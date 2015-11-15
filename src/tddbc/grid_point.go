package tddbc

import(
    "fmt"
)

type GridPoint struct {
    X int
    Y int
}

func (gridPoint GridPoint) GetNotation() string {
    return fmt.Sprintf("(%v,%v)", gridPoint.X, gridPoint.Y)
}

func (gridPoint GridPoint) IsNeighborOf(other GridPoint) bool {
    return gridPoint.isTopOf(other) ||
           gridPoint.isBottomOf(other) ||
           gridPoint.isRightOf(other) ||
           gridPoint.isLeftOf(other)
}

func (gridPoint GridPoint) isTopOf(other GridPoint) bool {
    return gridPoint.X == other.X && gridPoint.Y == other.Y + 1
}
func (gridPoint GridPoint) isBottomOf(other GridPoint) bool {
    return gridPoint.X == other.X && gridPoint.Y == other.Y - 1
}
func (gridPoint GridPoint) isRightOf(other GridPoint) bool {
    return gridPoint.X == other.X + 1 && gridPoint.Y == other.Y
}
func (gridPoint GridPoint) isLeftOf(other GridPoint) bool {
    return gridPoint.X == other.X - 1 && gridPoint.Y == other.Y
}
