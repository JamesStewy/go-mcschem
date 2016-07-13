package mcschem

import (
	"github.com/jteeuwen/mctools/anvil"
)

type Material string

const (
	Alpha   Material = "Alpha"
	Classic Material = "Classic"
)

type Block struct {
	Id   uint8
	Data uint8
}

type Schem struct {
	Width        uint16 // x
	Height       uint16 // y
	Length       uint16 // z
	WEOffsetX    int32
	WEOffsetY    int32
	WEOffsetZ    int32
	Materials    Material
	Blocks       []uint8
	Data         []uint8
	Entities     []*anvil.Entity
	TileEntities []*anvil.TileEntity
}

func (s *Schem) index(x, y, z int) int {
	if x < 0 || x >= int(s.Width) {
		panic("mcschem: x index out of range")
	}
	if y < 0 || y >= int(s.Height) {
		panic("mcschem: y index out of range")
	}
	if z < 0 || z >= int(s.Length) {
		panic("mcschem: z index out of range")
	}
	return (y*int(s.Length)+z)*int(s.Width) + x
}
