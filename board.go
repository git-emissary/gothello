package main

import (
	"fmt"
)

type Board struct {
	Length int
	disks  [][]rune
}

func NewEmptyBoard(length int) Board {
	disks := make([][]rune, length)
	for i := 0; i < length; i++ {
		disks[i] = make([]rune, length)
		for j := 0; j < length; j++ {
			disks[i][j] = BLANK
		}
	}

	return Board{
		Length: length,
		disks:  disks,
	}
}

func (b *Board) OutOfBounds(i int, j int) bool {
	return i < 0 || i >= b.Length || j < 0 || j > b.Length
}

func (b *Board) ToString() string {
	out := ""
	for i := 0; i < b.Length; i++ {
		out += "\n"
		for j := 0; j < b.Length; j++ {
			if j > 0 {
				out += " "
			}
			out += fmt.Sprintf("%c", b.disks[i][j])
		}
	}
	return out
}

// For debugging purposes.
func (b *Board) Print() {
	fmt.Print(b.ToString())
}

func (b *Board) setDisk(i int, j int, target rune) error {
	if b.OutOfBounds(i, j) {
		return ErrOutOfBounds
	}

	b.disks[i][j] = target
	return nil
}

// Sets all disks between positions (i1, j1) and (i2, j2) to target.
// Only straight directions are considered.
// Does not matter what disks are between those locations.
//
// You must use setDisk() method defined above.
//
// Hints: you may use Sign() and StartIndex() defined in utils.go.
//
// Eg. ROW
// From positions (2,0) to (2,3)
// _ _ _ _     _ _ _ _
// _ _ _ _     _ _ _ _
// x x x x  -> o o o o
// _ _ _ _     _ _ _ _
//
// Eg. COL
// From positions (0,1) to (3,1)
// _ x _ _     _ o _ _
// _ _ _ _     _ o _ _
// _ x _ _  -> _ o _ _
// _ o _ _     _ o _ _
//
// Eg. DIAGONAL (only square diagonal is considered)
// From positions (3,1) to (1,3)
// _ _ _ _     _ _ _ _
// _ _ _ x     _ _ _ o
// _ _ x _ ->  _ _ o _
// _ x _ _     _ o _ _
//
func (b *Board) Flip(i1 int, j1 int, i2 int, j2 int, target rune) error {
	deltaRow := i2 - i1
	deltaCol := j2 - j1

	if b.OutOfBounds(i1, j1) || b.OutOfBounds(i2, j2) {
		return ErrOutOfBounds
	}

	// Not proper diagonal.
	if deltaRow != 0 && deltaCol != 0 && IntAbs(deltaRow) != IntAbs(deltaCol) {
		return ErrNotSqDiagonal
	}

	// TODO: fill in rest of implementation
	return nil
}

// This method applies Flip() from position (i,j) to the closest target disk(s)
// in ALL directions. (UP, DOWN, LEFT, RIGHT, DIAGONALS)
// At (i,j) the disk is BLANK (already checked for you).
//
// NOTE: you are to reuse the Flip() method defined above.
//
// Eg.
// 5x5 input board, placing disk at 2,2 (centre of board).
// _ _ o _ _
// _ x _ _ _
// o x _ x x
// _ _ x x _
// _ _ o _ o
//
// Expected output board
// _ _ o _ _
// _ x _ _ _
// o o O x x
// _ _ o o _
// _ _ o _ o
//
func (b *Board) Cover(i int, j int, target rune) error {
	if b.OutOfBounds(i, j) {
		return ErrOutOfBounds
	}

	// You can only place a disk in an empty spot.
	if b.disks[i][j] != BLANK {
		return ErrInvalidPlacement
	}

	// TODO: fill in rest of implementation
	return nil
}
