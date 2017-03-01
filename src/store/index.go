package store

import (
	"github.com/hit9/htree"
	"github.com/edsrzf/mmap-go"
	"os"
)
type DefaultStore struct {
	tree htree.HTree
	mmap mmap.MMap
	file *os.File

}
type StoreConfig struct {
	FileMax uint32


}

func (ds *DefaultStore)Open(config StoreConfig)error{
	// 添加文件锁
	// 打开文件
	// 获取内存
	// mmap数据
	// 加载索引
	// 启动定时dump索引的协程



}
func (ds *DefaultStore)Close()error{

}
func (ds *DefaultStore)Put(key,value []byte)error{

}
func (ds *DefaultStore)Get(key []byte)(value []byte,err error){

}
func (ds *DefaultStore)Delete(key []byte)error{

}
func (ds *DefaultStore)Merge(){

}
func (ds *DefaultStore)DumpIndexFile()error{

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
