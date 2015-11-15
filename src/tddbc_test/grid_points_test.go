package tddbc_test

import (
    "testing"
    . "tddbc"
)

func Test_格子点の集合に指定した格子点が含まれているか判定できる(t *testing.T) {
    gridPoints := NewGridPoints( GridPoint{ X: 3, Y: 7 }, GridPoint{ X: 4, Y: 8 } )

    parameters := []struct {
        gridPoints GridPoints
        target GridPoint
        result bool
    }{
        { gridPoints: gridPoints, target: GridPoint{ X: 3, Y: 7 }, result: true },
        { gridPoints: gridPoints, target: GridPoint{ X: 4, Y: 8 }, result: true },
        { gridPoints: gridPoints, target: GridPoint{ X: 1, Y: 2 }, result: false } }

    for _, parameter := range parameters {
        actual := parameter.gridPoints.Contains(parameter.target)
        expected := parameter.result
        if expected != actual {
            t.Errorf("expected: %v, actual: %v", expected, actual)
        }
    }
}

func Test_格子点集合の格子点同士が連結しているか判定できる(t *testing.T) {
    parameters := []struct {
        gridPoints GridPoints
        result bool
    }{
        { gridPoints: NewGridPoints( GridPoint{ X: 3, Y: 7 }, GridPoint{ X: 3, Y: 8 } ), result: true },
        { gridPoints: NewGridPoints( GridPoint{ X: 3, Y: 7 }, GridPoint{ X: 4, Y: 8 } ), result: false },
        { gridPoints: NewGridPoints( GridPoint{ X: 3, Y: 7 }, GridPoint{ X: 3, Y: 8 }, GridPoint{ X: 4, Y: 8 } ), result: true },
        { gridPoints: NewGridPoints( GridPoint{ X: 3, Y: 7 }, GridPoint{ X: 4, Y: 8 }, GridPoint{ X: 5, Y: 8 } ), result: false },
    }

    for _, parameter := range parameters {
        actual := parameter.gridPoints.IsConnected()
        expected := parameter.result
        if expected != actual {
            t.Errorf("expected: %v, actual: %v", expected, actual)
        }
    }
}
