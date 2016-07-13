package mcschem

import (
	"compress/gzip"
	"github.com/jteeuwen/mctools/anvil"
	"github.com/jteeuwen/mctools/anvil/nbt"
	"io"
)

// Creates a new Schematic.
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

// Sets the block at the provided co-ordinates.
// Panics if an ordinate is out of range.
func (s *Schem) SetBlock(x, y, z int, b Block) {
	i := s.index(x, y, z)
	s.Blocks[i] = b.Id
	s.Data[i] = b.Data
}

// Marshals schematic to NBT format.
// Set gz=true to use gzip compression.
func (s *Schem) Marshal(w io.Writer, gz bool) error {
	if gz {
		wgz := gzip.NewWriter(w)
		defer wgz.Close()
		return nbt.Marshal(wgz, s)
	}
	return nbt.Marshal(w, s)
}
