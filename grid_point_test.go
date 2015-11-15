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

func Test_格子点の文字列表現を取得できる(t *testing.T) {
    gridPoint := GridPoint{ X: 1, Y: 2 }

    if "(1,2)" != gridPoint.getNotation() {
        t.Errorf("expected: \"(1,2)\", actual: \"%v\"", gridPoint.getNotation())
    }
}

func Test_格子点の文字列表現を取得できる2(t *testing.T) {
    gridPoint := GridPoint{ X: 4, Y: 7 }

    if "(4,7)" != gridPoint.getNotation() {
        t.Errorf("expected: \"(4,7)\", actual: \"%v\"", gridPoint.getNotation())
    }
}
