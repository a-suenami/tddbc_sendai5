package tddbc_test

import (
    "testing"
    . "tddbc"
)

type gridPointsContainsTestParameter struct {
    gridPoints GridPoints
    target GridPoint
    result bool
}

func Test_格子点の集合に指定した格子点が含まれているか判定できる(t *testing.T) {
    gridPoints := NewGridPoints( GridPoint{ X: 3, Y: 7 }, GridPoint{ X: 4, Y: 8 } )

    gridPointsContainsTestParameters := []gridPointsContainsTestParameter{
        gridPointsContainsTestParameter{ gridPoints: gridPoints, target: GridPoint{ X: 3, Y: 7 }, result: true },
        gridPointsContainsTestParameter{ gridPoints: gridPoints, target: GridPoint{ X: 4, Y: 8 }, result: true },
        gridPointsContainsTestParameter{ gridPoints: gridPoints, target: GridPoint{ X: 1, Y: 2 }, result: false } }

    for _, parameter := range gridPointsContainsTestParameters {
        actual := parameter.gridPoints.Contains(parameter.target)
        expected := parameter.result
        if expected != actual {
            t.Errorf("expected: %v, actual: %v", expected, actual)
        }
    }
}
