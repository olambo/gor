package main

import (
	"fmt"

	"gorgonia.org/gorgonia"
)

type gorXYZ struct {
	x *gorgonia.Node
	y *gorgonia.Node
	z *gorgonia.Node
	exprGraph *gorgonia.ExprGraph
}

func gorXplusYGraph() (gorXYZ, error) {
	creScalarNode := func(g *gorgonia.ExprGraph, name string) *gorgonia.Node {
		return gorgonia.NewScalar(g, gorgonia.Float32, gorgonia.WithName(name))
	}
	g := gorXYZ{}
	g.exprGraph = gorgonia.NewGraph()
	g.x = creScalarNode(g.exprGraph, "x")
	g.y = creScalarNode(g.exprGraph, "y")
	z, err := gorgonia.Add(g.x, g.y)
	if err != nil {
		return gorXYZ{}, fmt.Errorf("cannot create expression: %w", err)
	}
	g.z = z
	return g, nil
}

func gorXplusY(g gorXYZ, xval, yval float32) (float32, error) {
	var machine = gorgonia.NewTapeMachine(g.exprGraph)
	defer machine.Close()

	gorgonia.Let(g.x, xval)
	gorgonia.Let(g.y, yval)
	if err := machine.RunAll(); err != nil {
		return 0, fmt.Errorf("cannot run expression: %w", err)
	}
	ret, ok := g.z.Value().Data().(float32)
	if !ok {
		return 0, fmt.Errorf("cannot create float32 from : %v", g.z.Value())
	}
	return ret, nil
}

func main() {
	fmt.Println("Gorgonia stuff")
}
