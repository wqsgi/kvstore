package store

import (
	"github.com/hit9/htree"
	"github.com/edsrzf/mmap-go"
	"os"
)
type DefaultStore struct {
	tree htree.HTree
	mmap mmap.MMap

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
