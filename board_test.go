package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func assertEqualBoards(t *testing.T, b1 Board, b2 Board) {
	if !reflect.DeepEqual(b1, b2) {
		err := fmt.Errorf("\nwant: %s\n\ngot: %s", b1.ToString(), b2.ToString())
		t.Error(err)
    t.FailNow()
	}
}

func TestFlipRow(t *testing.T) {
	// _ _ _ _
	// _ _ _ _
	// _ _ _ _
	// _ _ _ _
	board := NewEmptyBoard(4)
	// _ _ _ _
	// o o o o
	// _ _ _ _
	// _ _ _ _
	expected := NewEmptyBoard(4)
	for j := 0; j < 4; j++ {
		expected.setDisk(1, j, ORB)
	}
  
	board.Flip(1, 0, 1, 3, ORB) // flip left to right
	assertEqualBoards(t, expected, board)
  
  board = NewEmptyBoard(4)
	board.Flip(1, 3, 1, 0, ORB) // flip right to left
	assertEqualBoards(t, expected, board)
}

func TestFlipCol(t *testing.T) {
	// _ _ _ _
	// _ _ _ _
	// _ _ _ _
	// _ _ _ _
	board := NewEmptyBoard(4)
	// _ _ _ o
	// _ _ _ o
	// _ _ _ o
	// _ _ _ o
	expected := NewEmptyBoard(4)
	for i := 0; i < 4; i++ {
		expected.setDisk(i, 3, ORB)
	}

	board.Flip(0, 3, 3, 3, ORB) // flip top to bottom
	assertEqualBoards(t, expected, board)

	board = NewEmptyBoard(4)
	board.Flip(3, 3, 0, 3, ORB) // flip bottom to top
	assertEqualBoards(t, expected, board)
}

// Minor diagonal is '/'
func TestMinorDiagonal(t *testing.T) {
	// _ _ _ _ _
	// _ _ _ _ _
	// _ _ _ _ _
	// _ _ _ _ _
	// _ _ _ _ _
	board := NewEmptyBoard(5)
	// _ _ _ o _
	// _ _ o _ _
	// _ o _ _ _
	// o _ _ _ _
	// _ _ _ _ _
	expected := NewEmptyBoard(5)
	for i, j := 3, 0; i >= 0; i, j = i-1, j+1 {
		expected.setDisk(i, j, ORB)
	}

	board.Flip(3, 0, 0, 3, ORB) // bottom to top
	assertEqualBoards(t, expected, board)

	board = NewEmptyBoard(5)
	board.Flip(0, 3, 3, 0, ORB) // top to bottom
	assertEqualBoards(t, expected, board)
}

// Major diagonal is '\'
func TestMajorDiagonal(t *testing.T) {
	// _ _ _ _ _
	// _ _ _ _ _
	// _ _ _ _ _
	// _ _ _ _ _
	// _ _ _ _ _
	board := NewEmptyBoard(5)
	// _ _ o _ _
	// _ _ _ o _
	// _ _ _ _ o
	// _ _ _ _ _
	// _ _ _ _ _
	expected := NewEmptyBoard(5)
	for i, j := 0, 2; j < 5; i, j = i+1, j+1 {
		expected.setDisk(i, j, ORB)
	}

	board.Flip(0, 2, 2, 4, ORB) // top to bottom
	assertEqualBoards(t, expected, board)

	board = NewEmptyBoard(5)
	board.Flip(2, 4, 0, 2, ORB) // bottom to top
	assertEqualBoards(t, expected, board)
}

func TestCannotCoverNonEmptySpot(t *testing.T) {
	board := NewEmptyBoard(2)
	board.setDisk(0, 0, ORB)

	err := board.Cover(0, 0, CROSS)
	want := ErrInvalidPlacement
	if !errors.Is(err, want) {
		t.Errorf("Expected \"%s\". Got %+v.", want, err)
	}
}

func TestCover(t *testing.T) {
	// _ _ _ o _ _ _ _ (0)
	// _ _ _ x _ _ o _ (1)
	// _ o _ x _ x _ _ (2)
	// _ _ x x _ o _ _ (3)
	// o x x _ x x x _ (4)
	// _ _ _ o x _ _ _ (5)
	// _ o o o x x _ _ (6)
	// _ _ _ x _ _ o _ (7)
	board := NewEmptyBoard(8)
	// 1st row
	board.setDisk(0, 3, ORB)
	// 2nd row
	board.setDisk(1, 3, CROSS)
	board.setDisk(1, 6, ORB)
	// 3rd row
	board.setDisk(2, 1, ORB)
	board.setDisk(2, 3, CROSS)
	board.setDisk(2, 5, CROSS)
	// 4th row
	board.setDisk(3, 2, CROSS)
	board.setDisk(3, 3, CROSS)
	board.setDisk(3, 5, ORB)
	// 5th row
	board.setDisk(4, 0, ORB)
	board.setDisk(4, 1, CROSS)
	board.setDisk(4, 2, CROSS)
	board.setDisk(4, 4, CROSS)
	board.setDisk(4, 5, CROSS)
	board.setDisk(4, 6, CROSS)
	// 6th row
	board.setDisk(5, 3, ORB)
	board.setDisk(5, 4, CROSS)
	// 7th row
	board.setDisk(6, 1, ORB)
	board.setDisk(6, 2, ORB)
	board.setDisk(6, 3, ORB)
	board.setDisk(6, 4, CROSS)
	board.setDisk(6, 5, CROSS)
	// 8th row
	board.setDisk(7, 3, CROSS)
	board.setDisk(7, 6, ORB)

	// _ _ _ o _ _ _ _
	// _ _ _ o _ _ o _
	// _ o _ o _ x _ _
	// _ _ o o _ o _ _
	// o o o O x x x _
	// _ _ _ o o _ _ _
	// _ o o o x o _ _
	// _ _ _ x _ _ o _
	expected := NewEmptyBoard(8)
	// 1st row
	expected.setDisk(0, 3, ORB)
	// 2nd row
	expected.setDisk(1, 3, ORB)
	expected.setDisk(1, 6, ORB)
	// 3rd row
	expected.setDisk(2, 1, ORB)
	expected.setDisk(2, 3, ORB)
	expected.setDisk(2, 5, CROSS)
	// 4th row
	expected.setDisk(3, 2, ORB)
	expected.setDisk(3, 3, ORB)
	expected.setDisk(3, 5, ORB)
	// 5th row
	expected.setDisk(4, 0, ORB)
	expected.setDisk(4, 1, ORB)
	expected.setDisk(4, 2, ORB)
	expected.setDisk(4, 3, ORB)
	expected.setDisk(4, 4, CROSS)
	expected.setDisk(4, 5, CROSS)
	expected.setDisk(4, 6, CROSS)
	// 6th row
	expected.setDisk(5, 3, ORB)
	expected.setDisk(5, 4, ORB)
	// 7th row
	expected.setDisk(6, 1, ORB)
	expected.setDisk(6, 2, ORB)
	expected.setDisk(6, 3, ORB)
	expected.setDisk(6, 4, CROSS)
	expected.setDisk(6, 5, ORB)
	// 8th row
	expected.setDisk(7, 3, CROSS)
	expected.setDisk(7, 6, ORB)

	board.Cover(4, 3, ORB)
	assertEqualBoards(t, expected, board)
}
