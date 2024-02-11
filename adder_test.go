run package tests | run file tests
package main

import "testing"

run test|debug test
func TestAdder(t *testing.T) {
    result := adder(1, 2)
    if result != 3 {
        t.Errorf("Expected %d but got %d", 3, result)
    }
}
