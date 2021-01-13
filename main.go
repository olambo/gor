package main

import (
	"fmt"

	"gorgonia.org/gorgonia"
)

func gorXplusY(xval, yval float64) (float64, error) {
	creScalarNode := func(g *gorgonia.ExprGraph, name string) *gorgonia.Node {
		return gorgonia.NewScalar(g, gorgonia.Float32, gorgonia.WithName(name))
	}
	exprGraph := gorgonia.NewGraph()
	x := creScalarNode(exprGraph, "x")
	y := creScalarNode(exprGraph, "y")
	z, err := gorgonia.Add(x, y)
	if err != nil {
		return 0, fmt.Errorf("cannot create expression: %w", err)
	}

	machine := gorgonia.NewTapeMachine(exprGraph)
	defer machine.Close()

	gorgonia.Let(x, xval)
	gorgonia.Let(y, yval)
	if err = machine.RunAll(); err != nil {
		return 0, fmt.Errorf("cannot run expression: %w", err)
	}
	ret, ok := z.Value().Data().(float64)
	if !ok {
		return 0, fmt.Errorf("cannot create float64 from : %v", z.Value())
	}
	return ret, nil
}

func main() {
	fmt.Println("Gorgonia stuff")
}
