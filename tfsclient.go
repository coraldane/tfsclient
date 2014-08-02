package tfsclient

// #cgo CFLAGS: -I/usr/local/tfs-2.2/include
// #cgo LDFLAGS: -L/usr/local/tfs-2.2/lib -ltfsclient_c -ltbsys -ltbnet -lz
// #include <stdlib.h>
// #include <stdint.h>
// #include "tfs_client_capi.h"
import "C"
import (
	"errors"
	"unsafe"
)

var (
	ErrInitTfsClient = errors.New("init tfs client error")
	ErrTfsError      = errors.New("tfs error")
)

type TfsClient struct {
}

func NewTfsClient(nsAddr string, cacheTime int32, cacheItems int32, startBg bool) (*TfsClient, error) {
	var ns_addr *C.char
	var ret, start_bg C.int
	var cache_time, cache_items C.int32_t

	if nsAddr != "" {
		ns_addr = C.CString(nsAddr)
		defer C.free(unsafe.Pointer(ns_addr))
	}
	if startBg {
		start_bg = 1
	} else {
		start_bg = 0
	}
	cache_time = C.int32_t(cacheTime)
	cache_items = C.int32_t(cacheItems)

	ret = C.t_initialize(ns_addr, cache_time, cache_items, start_bg)

	if int(ret) != TFS_SUCCESS {
		return nil, ErrInitTfsClient
	}
	return &TfsClient{}, nil
}

func (c *TfsClient) SetDefaultServer(nsAddr string, cacheTime, cacheItems int32) error {
	var ns_addr *C.char
	var ret C.int
	var cache_time, cache_items C.int32_t

	if nsAddr != "" {
		ns_addr = C.CString(nsAddr)
		defer C.free(unsafe.Pointer(ns_addr))
	}
	cache_time = C.int32_t(cacheTime)
	cache_items = C.int32_t(cacheItems)

	ret = C.t_set_default_server(ns_addr, cache_time, cache_items)

	if int(ret) != TFS_SUCCESS {
		return ErrInitTfsClient
	}
	return nil
}

func (c *TfsClient) Destroy() error {
	var ret C.int

	ret = C.t_destroy()
	if int(ret) != TFS_SUCCESS {
		return ErrTfsError
	}
	return nil
}

func (c *TfsClient) Open(filename, suffix string, flags int, localKey string) (int, error) {
	var c_filename *C.char
	var c_suffix *C.char
	var local_key *C.char
	var fd C.int

	if filename != "" {
		c_filename = C.CString(filename)
		defer C.free(unsafe.Pointer(c_filename))
	}
	if suffix != "" {
		c_suffix = C.CString(suffix)
		defer C.free(unsafe.Pointer(c_suffix))
	}
	if localKey != "" {
		local_key = C.CString(localKey)
		defer C.free(unsafe.Pointer(local_key))
	}

	fd = C.t_open(c_filename, c_suffix, C.int(flags), local_key)

	if fd <= 0 {
		return 0, ErrTfsError
	}
	return int(fd), nil
}

func (c *TfsClient) Open2(filename, suffix, nsAddr string, flags int, localKey string) (int, error) {
	var c_filename *C.char
	var c_suffix *C.char
	var ns_addr *C.char
	var local_key *C.char
	var fd C.int

	if filename != "" {
		c_filename = C.CString(filename)
		defer C.free(unsafe.Pointer(c_filename))
	}
	if suffix != "" {
		c_suffix = C.CString(suffix)
		defer C.free(unsafe.Pointer(c_suffix))
	}
	if localKey != "" {
		local_key = C.CString(localKey)
		defer C.free(unsafe.Pointer(local_key))
	}
	if nsAddr != "" {
		ns_addr = C.CString(nsAddr)
		defer C.free(unsafe.Pointer(ns_addr))
	}

	fd = C.t_open2(c_filename, c_suffix, ns_addr, C.int(flags), local_key)

	if fd <= 0 {
		return 0, ErrTfsError
	}
	return int(fd), nil
}

func (c *TfsClient) Read(fd int, buf []byte) int64 {
	var count C.int64_t

	cbuf := unsafe.Pointer(&buf[0])
	count = C.int64_t(len(buf))

	ret := C.t_read(C.int(fd), cbuf, count)
	return int64(ret)
}

func (c *TfsClient) Read2(fd int, buf []byte) (TfsFileStat, int64) {
	var stat TfsFileStat
	var cstat C.TfsFileStat
	var count C.int64_t

	cbuf := unsafe.Pointer(&buf[0])
	count = C.int64_t(len(buf))

	ret := C.t_readv2(C.int(fd), cbuf, count, &cstat)

	stat.FileId = uint64(cstat.file_id_)
	stat.Offset = int32(cstat.offset_)
	stat.Size = int64(cstat.size_)
	stat.USize = int64(cstat.usize_)
	stat.ModifyTime = int64(cstat.modify_time_)
	stat.CreateTime = int64(cstat.create_time_)
	stat.Flag = int32(cstat.flag_)
	stat.CRC = uint32(cstat.crc_)
	return stat, int64(ret)
}

func (c *TfsClient) Write(fd int, buf []byte) int64 {
	var count C.int64_t

	cbuf := unsafe.Pointer(&buf[0])
	count = C.int64_t(len(buf))

	ret := C.t_write(C.int(fd), cbuf, count)
	return int64(ret)
}

func (c *TfsClient) Seek(fd int, offset int64, whence int) int64 {
	ret := C.t_lseek(C.int(fd), C.int64_t(offset), C.int(whence))
	return int64(ret)
}

func (c *TfsClient) Pread(fd int, buf []byte, offset int64) int64 {
	var count C.int64_t

	cbuf := unsafe.Pointer(&buf[0])
	count = C.int64_t(len(buf))

	ret := C.t_pread(C.int(fd), cbuf, count, C.int64_t(offset))
	return int64(ret)
}

func (c *TfsClient) Pwrite(fd int, buf []byte, offset int64) int64 {
	var count C.int64_t

	cbuf := unsafe.Pointer(&buf[0])
	count = C.int64_t(len(buf))

	ret := C.t_pwrite(C.int(fd), cbuf, count, C.int64_t(offset))
	return int64(ret)
}

func (c *TfsClient) GetFileLength(fd int) int64 {
	return int64(C.t_get_file_length(C.int(fd)))
}

func (c *TfsClient) Fstat(fd int, mode byte) (TfsFileStat, error) {
	var stat TfsFileStat
	var cstat C.TfsFileStat

	ret := C.t_fstat(C.int(fd), &cstat, C.TfsStatType(mode))
	if int(ret) != TFS_SUCCESS {
		return stat, ErrTfsError
	}
	stat.FileId = uint64(cstat.file_id_)
	stat.Offset = int32(cstat.offset_)
	stat.Size = int64(cstat.size_)
	stat.USize = int64(cstat.usize_)
	stat.ModifyTime = int64(cstat.modify_time_)
	stat.CreateTime = int64(cstat.create_time_)
	stat.Flag = int32(cstat.flag_)
	stat.CRC = uint32(cstat.crc_)
	return stat, nil
}

func (c *TfsClient) Close(fd int) (string, error) {
	filename := make([]byte, TFS_FILE_LEN)

	c_filename := (*C.char)(unsafe.Pointer(&filename[0]))
	ret := C.t_close(C.int(fd), c_filename, C.int32_t(TFS_FILE_LEN))

	if int(ret) != TFS_SUCCESS {
		return "", ErrTfsError
	}
	return string(filename[:FILE_NAME_LEN]), nil
}

func (c *TfsClient) Unlink(filename, suffix string, action int) (int64, error) {
	var c_filename, c_suffix *C.char
	c_filename = C.CString(filename)
	if suffix != "" {
		c_suffix = C.CString(suffix)
		defer C.free(unsafe.Pointer(c_suffix))
	}
	var filesize C.int64_t

	ret := C.t_unlink(c_filename, c_suffix, &filesize, C.TfsUnlinkType(action))

	C.free(unsafe.Pointer(c_filename))

	if int(ret) != TFS_SUCCESS {
		return 0, ErrTfsError
	}
	return int64(filesize), nil
}

func (c *TfsClient) Unlink2(filename, suffix string, action byte, nsAddr string) (int64, error) {
	var c_filename *C.char
	var c_suffix *C.char
	var ns_addr *C.char
	var filesize C.int64_t

	c_filename = C.CString(filename)
	if suffix != "" {
		c_suffix = C.CString(suffix)
		defer C.free(unsafe.Pointer(c_suffix))
	}
	if nsAddr != "" {
		ns_addr = C.CString(nsAddr)
		defer C.free(unsafe.Pointer(ns_addr))
	}

	ret := C.t_unlink2(c_filename, c_suffix, &filesize, C.TfsUnlinkType(action), ns_addr)

	C.free(unsafe.Pointer(c_filename))

	if int(ret) != TFS_SUCCESS {
		return 0, ErrTfsError
	}
	return int64(filesize), nil
}

func (c *TfsClient) SetOptionFlag(fd int, optionFlag byte) error {
	ret := C.t_set_option_flag(C.int(fd), C.OptionFlag(optionFlag))

	if int(ret) != TFS_SUCCESS {
		return ErrTfsError
	}
	return nil
}

func (c *TfsClient) SetCacheItems(val int64) {
	C.t_set_cache_items(C.int64_t(val))
}

func (c *TfsClient) GetCacheItems() int64 {
	return int64(C.t_get_cache_items())
}

func (c *TfsClient) SetCacheTime(val int64) {
	C.t_set_cache_time(C.int64_t(val))
}

func (c *TfsClient) GetCacheTime() int64 {
	return int64(C.t_get_cache_time())
}

func (c *TfsClient) SetSegmentSize(val int64) {
	C.t_set_segment_size(C.int64_t(val))
}

func (c *TfsClient) GetSegmentSize() int64 {
	return int64(C.t_get_segment_size())
}

func (c *TfsClient) SetBatchCount(val int64) {
	C.t_set_batch_count(C.int64_t(val))
}

func (c *TfsClient) GetBatchCount() int64 {
	return int64(C.t_get_batch_count())
}

func (c *TfsClient) SetStatInterVal(val int64) {
	C.t_set_stat_interval(C.int64_t(val))
}

func (c *TfsClient) GetStatInterval() int64 {
	return int64(C.t_get_stat_interval())
}

func (c *TfsClient) SetGcInterVal(val int64) {
	C.t_set_gc_interval(C.int64_t(val))
}

func (c *TfsClient) GetGcInterval() int64 {
	return int64(C.t_get_gc_interval())
}

func (c *TfsClient) SetGcExpiredTime(val int64) {
	C.t_set_gc_expired_time(C.int64_t(val))
}

func (c *TfsClient) GetGcExpiredTime() int64 {
	return int64(C.t_get_gc_expired_time())
}

func (c *TfsClient) SetBatchTimeout(val int64) {
	C.t_set_batch_timeout(C.int64_t(val))
}

func (c *TfsClient) GetBatchTimeout() int64 {
	return int64(C.t_get_batch_timeout())
}

func (c *TfsClient) SetWaitTimeout(val int64) {
	C.t_set_wait_timeout(C.int64_t(val))
}

func (c *TfsClient) GetWaitTimeout() int64 {
	return int64(C.t_get_wait_timeout())
}

func (c *TfsClient) SetClientRetryCount(val int64) {
	C.t_set_client_retry_count(C.int64_t(val))
}

func (c *TfsClient) GetClientRetryCount() int64 {
	return int64(C.t_get_client_retry_count())
}

func (c *TfsClient) SetLogLevel(level string) {
	if level == "" {
		return
	}
	clevel := C.CString(level)
	C.t_set_log_level(clevel)
	C.free(unsafe.Pointer(clevel))
}

func (c *TfsClient) SetLogFile(file string) {
	if file == "" {
		return
	}
	cfile := C.CString(file)
	C.t_set_log_file(cfile)
	C.free(unsafe.Pointer(cfile))

}

func (c *TfsClient) GetServerId() uint64 {
	return uint64(C.t_get_server_id())
}

func (c *TfsClient) GetClusterId() int32 {
	return int32(C.t_get_cluster_id())
}
