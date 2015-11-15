package tddbc_test

import (
    "testing"
    . "tddbc"
)

func Test_格子点の集合(t *testing.T) {
    gridPoints := NewGridPoints( GridPoint{ X: 3, Y: 7 }, GridPoint{ X: 4, Y: 8 } )

    target := GridPoint{ X: 3, Y: 7 }
    if gridPoints.Contains(target) == false {
        t.Fail()
    }
}
