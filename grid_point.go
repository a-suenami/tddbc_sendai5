package tddbc

import(
    "fmt"
)

type GridPoint struct {
    X int
    Y int
}

func (gridPoint GridPoint) getNotation() string {
    return fmt.Sprintf("(%v,%v)", gridPoint.X, gridPoint.Y)
}
