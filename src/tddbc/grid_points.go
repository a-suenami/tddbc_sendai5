package tddbc

type GridPoints struct {
    gridPoints []GridPoint
}

func NewGridPoints(gridPoint1 GridPoint, gridPoint2 GridPoint) GridPoints {
    return GridPoints{ gridPoints: []GridPoint{ gridPoint1, gridPoint2 } }
}

func (gridPoints GridPoints) Contains(gridPoint GridPoint) bool {
    for _, gp := range gridPoints.gridPoints {
        if gp.HasSameCoordinatesWith(gridPoint) {
            return true
        }
    }
    return false;
}
