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
