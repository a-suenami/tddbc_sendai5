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
    if len(gridPoints.gridPoints) == 3 {
        return (gridPoints.gridPoints[0].IsNeighborOf(gridPoints.gridPoints[1]) && gridPoints.gridPoints[0].IsNeighborOf(gridPoints.gridPoints[2])) ||
               (gridPoints.gridPoints[1].IsNeighborOf(gridPoints.gridPoints[2]) && gridPoints.gridPoints[1].IsNeighborOf(gridPoints.gridPoints[0])) ||
               (gridPoints.gridPoints[2].IsNeighborOf(gridPoints.gridPoints[0]) && gridPoints.gridPoints[2].IsNeighborOf(gridPoints.gridPoints[1]))
    } else {
        return gridPoints.gridPoints[0].IsNeighborOf(gridPoints.gridPoints[1])
    }
}
