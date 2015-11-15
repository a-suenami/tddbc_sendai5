package tddbc

import (
    "testing"
)

func Test_Xが正しく取得できる(t *testing.T) {
    gridPoint := GridPoint{ X: 1, Y: 2 }

    if 1 != gridPoint.X {
        t.Errorf("expected: 1, actual: %v", gridPoint.X)
    }
}

func Test_Yが正しく取得できる(t *testing.T) {
    gridPoint := GridPoint{ X: 1, Y: 2 }

    if 2 != gridPoint.Y {
        t.Errorf("expected: 2, actual: %v", gridPoint.Y)
    }
}

type gridPointTestParameter struct {
    gridPoint GridPoint
    notation string
}

func Test_格子点の文字列表現を取得できる(t *testing.T) {
    gridPointTestParameters := []gridPointTestParameter{
        gridPointTestParameter{ gridPoint: GridPoint{ X: 1, Y: 2 }, notation: "(1,2)" },
        gridPointTestParameter{ gridPoint: GridPoint{ X: 4, Y: 7 }, notation: "(4,7)" } }

    for _, parameter := range gridPointTestParameters {
        if parameter.notation != parameter.gridPoint.getNotation() {
            t.Errorf("expected: \"%v\", actual: \"%v\"", parameter.notation, parameter.gridPoint.getNotation())
        }
    }
}
