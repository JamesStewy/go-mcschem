/*
Package mcschem allows for the creation, marshalling and unmarshalling of Minecraft schematics.
Based of the Minecraft schematic format found here: http://minecraft.gamepedia.com/Schematic_file_format.

See test file for an example usage of mcschem.
*/
package mcschem

import (
	"github.com/jteeuwen/mctools/anvil"
)

// Type of world file.
// Either "Alpha" (Default) or "Classic".
type Material string

const (
	Alpha   Material = "Alpha"
	Classic Material = "Classic"
)

// Minecraft block.
type Block struct {
	// Block ID
	Id uint8

	// Block Data. Zero if no extra data is needed.
	Data uint8
}

type Schem struct {
	// Schematic Size
	Width  uint16 // x
	Height uint16 // y
	Length uint16 // z

	// WorldEdit Offset
	WEOffsetX int32
	WEOffsetY int32
	WEOffsetZ int32

	// Type of world file
	Materials Material

	// Block data
	Blocks []uint8
	Data   []uint8

	// Entity data. Uses definitions in "github.com/jteeuwen/mctools/anvil".
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
