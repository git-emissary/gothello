package main

import "fmt"

var ErrOutOfBounds = fmt.Errorf("Index out of bounds.")
var ErrNotSqDiagonal = fmt.Errorf("Not square diagonal.")
var ErrInvalidPlacement = fmt.Errorf("Cannot place disk at location.")
