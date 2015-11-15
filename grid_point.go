package tddbc

type GridPoint struct {
    X int
    Y int
}

func (GridPoint) getNotation() string {
    return "(1,2)"
}
