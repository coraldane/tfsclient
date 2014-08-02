package tfsclient

/* error code */
const (
	TFS_SUCCESS          = 0
	TFS_ERROR            = 1
	EXIT_INVALIDFD_ERROR = -1005
)

/* block cache default config */
const (
	DEFAULT_BLOCK_CACHE_TIME  = 1800
	DEFAULT_BLOCK_CACHE_ITEMS = 500000
)

/* cluster group count and group seq default value */
const (
	DEFAULT_CLUSTER_GROUP_COUNT = 1
	DEFAULT_CLUSTER_GROUP_SEQ   = 0
)

/* tfs file name standard name length */
const (
	FILE_NAME_LEN               = 18
	TFS_FILE_LEN                = FILE_NAME_LEN + 1
	FILE_NAME_EXCEPT_SUFFIX_LEN = 12
	MAX_FILE_NAME_LEN           = 128
	MAX_SUFFIX_LEN              = MAX_FILE_NAME_LEN - TFS_FILE_LEN
	STANDARD_SUFFIX_LEN         = 4
)

type TfsFileStat struct {
	FileId      uint64
	Offset      int32
	Size, USize int64
	ModifyTime  int64
	CreateTime  int64
	Flag        int32
	CRC         uint32
}

/* open flag */
const (
	T_DEFAULT = 0
	T_READ    = 1 << (iota - 1)
	T_WRITE
	T_CREATE
	T_NEWBLK
	T_NOLEASE
	T_STAT
	T_LARGE
	T_UNLINK
	T_FORCE
)

/* tfs seek type */
const (
	T_SEEK_SET = iota
	T_SEEK_CUR
	T_SEEK_END
)

/* tfs stat type */
const (
	NORMAL_STAT = iota
	FORCE_STAT
)

/* unlink type */
const (
	DELETE   = 0
	UNDELETE = 2
	CONCEAL  = 4
	REVEAL   = 6
	OVERRIDE = 128
)

/* option flag */
const (
	TFS_FILE_DEFAULT_OPTION = iota
	TFS_FILE_NO_SYNC_LOG
	TFS_FILE_CLOSE_FLAG_WRITE_DATA_FAILED
)

/* read data option flag */
const (
	READ_DATA_OPTION_FLAG_NORMAL = iota
	READ_DATA_OPTION_FLAG_FORCE
)

/* cache type */
const (
	LOCAL_CACHE = iota
	REMOTE_CACHE
)

/* use cache flag */
const (
	USE_CACHE_FLAG_NO     = 0x00
	USE_CACHE_FLAG_LOCAL  = 0x01
	USE_CACHE_FLAG_REMOTE = 0x02
)
