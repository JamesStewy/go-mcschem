package mcschem

import (
	"compress/gzip"
	"github.com/jteeuwen/mctools/anvil"
	"github.com/jteeuwen/mctools/anvil/nbt"
	"io"
)

func New(width, length, height int) *Schem {
	return &Schem{
		Width:        uint16(width),
		Height:       uint16(height),
		Length:       uint16(length),
		Materials:    Alpha,
		Blocks:       make([]uint8, width*length*height),
		Data:         make([]uint8, width*length*height),
		Entities:     make([]*anvil.Entity, 0),
		TileEntities: make([]*anvil.TileEntity, 0),
	}
}

func (s *Schem) SetBlock(x, y, z int, b Block) {
	i := s.index(x, y, z)
	s.Blocks[i] = b.Id
	s.Data[i] = b.Data
}

func (s *Schem) Marshal(w io.Writer, gz bool) error {
	if gz {
		wgz := gzip.NewWriter(w)
		defer wgz.Close()
		return nbt.Marshal(wgz, s)
	}
	return nbt.Marshal(w, s)
}
