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

func Test_上の格子点を隣り合っていると判定する(t *testing.T) {
    gridPoint := GridPoint{ X: 3, Y: 7 }
    target := GridPoint{ X: 3, Y: 8 }

    if gridPoint.isNeighborOf(target) == false {
        t.Fail()
    }
}

func Test_右の格子点を隣り合っていると判定する(t *testing.T) {
    gridPoint := GridPoint{ X: 3, Y: 7 }
    target := GridPoint{ X: 4, Y: 7 }

    if gridPoint.isNeighborOf(target) == false {
        t.Fail()
    }
}

func Test_下の格子点を隣り合っていると判定する(t *testing.T) {
    gridPoint := GridPoint{ X: 3, Y: 7 }
    target := GridPoint{ X: 3, Y: 6 }

    if gridPoint.isNeighborOf(target) == false {
        t.Fail()
    }
}

func Test_左の格子点を隣り合っていると判定する(t *testing.T) {
    gridPoint := GridPoint{ X: 3, Y: 7 }
    target := GridPoint{ X: 2, Y: 7 }

    if gridPoint.isNeighborOf(target) == false {
        t.Fail()
    }
}

func Test_右上の格子点を隣り合っていないと判定する(t *testing.T) {
    gridPoint := GridPoint{ X: 3, Y: 7 }
    target := GridPoint{ X: 4, Y: 8 }

    if gridPoint.isNeighborOf(target) {
        t.Fail()
    }
}

func Test_右下の格子点を隣り合っていないと判定する(t *testing.T) {
    gridPoint := GridPoint{ X: 3, Y: 7 }
    target := GridPoint{ X: 4, Y: 6 }

    if gridPoint.isNeighborOf(target) {
        t.Fail()
    }
}

func Test_左上の格子点を隣り合っていないと判定する(t *testing.T) {
    gridPoint := GridPoint{ X: 3, Y: 7 }
    target := GridPoint{ X: 2, Y: 8 }

    if gridPoint.isNeighborOf(target) {
        t.Fail()
    }
}

func Test_左下の格子点を隣り合っていないと判定する(t *testing.T) {
    gridPoint := GridPoint{ X: 3, Y: 7 }
    target := GridPoint{ X: 2, Y: 6 }

    if gridPoint.isNeighborOf(target) {
        t.Fail()
    }
}
