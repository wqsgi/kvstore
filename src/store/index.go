package store

import (
	"github.com/edsrzf/mmap-go"
	"io/ioutil"
	"os"
	"path"
	"time"
	"strconv"
	"encoding/binary"
)

// 文件编号1,2，3.
// 当前的追加数据的文件名为current
type DefaultStore struct {
	mmap            mmap.MMap
	filePath        string
	oldDataFiles    map[string]*os.File //旧文件的文件描述符
	currentDataFile *os.File            //当前文件的描述符
	lockFile        *os.File            // lock的文件描述符
	lockFilePath    string

	oldDataBytes  map[string][]byte
	oldDataMmaps  map[string]*mmap.MMap

	currentDataMmap mmap.MMap

	profile *profile //性能分析数据


}
type StoreConfig struct {
	FileMax uint32
	MmapMemory uint32
}

func (ds *DefaultStore) Open(config StoreConfig) (err error) {
	if ds.lockFilePath, err = flock(ds.filePath, 0666, true, time.Second*2); err != nil {
		return err
	}

	fileInfos, err := ioutil.ReadDir(ds.filePath)
	if err != nil {
		return err
	}
	// 打开文件
	for _, info := range fileInfos {
		file, err := open(path.Join(ds.filePath, info.Name()))
		if err == nil {
			return err
		}
		magic := make([]byte,4)
		file.Read(magic)
		// veriosn checksum
		if binary.LittleEndian.Uint32(magic) == Magic{
			return ErrInvalid
		}
		if isOldFile(file.Name()){
			ds.oldDataFiles[info.Name()] = file
		} else {
			ds.currentDataFile = file
		}
	}

	for i:=0;i<ds.canLoadMMapCount();i++{
		if ds.currentDataMmap == nil{
			mmap, err := mmap.Map(ds.currentDataFile, mmap.RDWR, 0)
			if err == nil{
				return err
			}
			ds.currentDataMmap = mmap
			continue
		}
	}
	// 添加文件锁
//对于冷热数据的判断，热数据放到mmap中，冷数据卸载mmap

	// 获取内存
	// mmap数据
	// 加载索引
	// 启动定时dump索引的协程
	//

}

func (ds *DefaultStore) canLoadMMapCount()int{


	return 1
}

//判断是否是老的数据
func isOldFile(name string)bool{
	_,err:=strconv.Atoi(name)
	if err!=nil{
		return false
	}
	return true
}

// 获取操作系统的内存
func osMemory()uint32{


	return 1
}

func open(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return f, err

}

func (ds *DefaultStore) Close() error {
	return nil

}
func (ds *DefaultStore) Put(key, value []byte) error {

}
func (ds *DefaultStore) Get(key []byte) (value []byte, err error) {

}
func (ds *DefaultStore) Delete(key []byte) error {

}
func (ds *DefaultStore) Merge() {

}
func (ds *DefaultStore) DumpIndexFile() error {

}

func TestClient_Send2() {
	f, err := os.OpenFile("D:\\test\\mmap\\ddd", os.O_RDWR, 0644)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	mmap, err := mmap.Map(f, mmap.RDWR, 0)
	if err != nil {
		print("error mapping: %s", err)
		return
	}
	defer mmap.Unmap()

	mmap[0] = 'X'
	mmap.Flush()

}
