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

func (gridPoint GridPoint) isNeighborOf(other GridPoint) bool {
    return (gridPoint.X == other.X     && gridPoint.Y == other.Y - 1) ||
           (gridPoint.X == other.X - 1 && gridPoint.Y == other.Y    ) ||
           (gridPoint.X == other.X     && gridPoint.Y == other.Y + 1) ||
           (gridPoint.X == other.X + 1 && gridPoint.Y == other.Y    )
}
