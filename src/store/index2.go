package store

import (
	"github.com/hit9/htree"
	"os"
)

type indexStore struct {
	tree *htree.HTree
	file *os.File
}

func (is *indexStore) load(file *os.File){

	is.tree.Put()

}

type Offset struct {
	fid uint16
	key []byte
	offset uint32
	size uint16
	version uint32
}


