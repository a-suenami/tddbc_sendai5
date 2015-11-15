package tddbc

type GridPoints struct {
    gridPoints []GridPoint
}

func NewGridPoints(gridPoints ...GridPoint) GridPoints {
    return GridPoints{ gridPoints: gridPoints }
}

func (gridPoints GridPoints) Contains(gridPoint GridPoint) bool {
    for _, gp := range gridPoints.gridPoints {
        if gp.HasSameCoordinatesWith(gridPoint) {
            return true
        }
    }
    return false
}

func (gridPoints GridPoints) IsConnected() bool {
    return gridPoints.gridPoints[0].IsNeighborOf(gridPoints.gridPoints[1])
}
