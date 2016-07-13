package mcschem

import (
	"compress/gzip"
	"github.com/jteeuwen/mctools/anvil/nbt"
	"io"
)

func Unmarshal(r io.Reader, gz bool) (schem *Schem, err error) {
	if gz {
		if r, err = gzip.NewReader(r); err != nil {
			return
		}
	}
	schem = &Schem{}
	err = nbt.Unmarshal(r, schem)
	return
}

func (s *Schem) Block(x, y, z int) Block {
	i := s.index(x, y, z)
	return Block{
		Id:   s.Blocks[i],
		Data: s.Data[i],
	}
}
