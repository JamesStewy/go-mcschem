package mcschem

import (
	"compress/gzip"
	"github.com/jteeuwen/mctools/anvil/nbt"
	"io"
)

// Unmarshals schematic from NBT format.
// Set gz=true if the Schematic uses gzip compression.
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

// Returns the block at the provided co-ordinates.
// Panics if an ordinate is out of range.
func (s *Schem) Block(x, y, z int) Block {
	i := s.index(x, y, z)
	return Block{
		Id:   s.Blocks[i],
		Data: s.Data[i],
	}
}
