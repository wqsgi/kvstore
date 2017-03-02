package store


// 日志文件
type HLog struct {
	crc32 int32 // 4byte
	timestamp int64 //8byte
	keySize uint16 // 2byte
	valueSize uint16 // 2byte
	key []byte
	value []byte
	version uint32 //4byte

}


//索引文件
type HIndexFile struct {

}

// 存储接口
type Store interface {
	Open(config StoreConfig)error
	Close()error
	Put(key,value []byte)error
	Get(key []byte)(value []byte,error)
	Delete(key []byte)error
	Merge()// 合并log数据。
	DumpIndexFile()error // dump索引数据。

}



