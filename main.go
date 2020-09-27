package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	xys, err := readData("data.txt")
	if err != nil {
		log.Fatalf("could not read data.txt: %v", err)
	}

	err = plotData("out.png", xys)
	if err != nil {
		log.Fatalf("could not plot : %v", err)
	}
}

func readData(path string) (plotter.XYs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var xys plotter.XYs
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y float64
		_, err = fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}
		xys = append(xys, struct{ X, Y float64 }{x, y})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("cound not scan: %v", err)
	}
	return xys, nil
}

func plotData(path string, xys plotter.XYs) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}
	defer f.Close()

	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("could not create plot %v", err)
	}

	// create scatter with all data points
	s, err := plotter.NewScatter(xys)
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	var x, c float64
	x = 1
	c = 1
	// create fake linear regression
	l, err := plotter.NewLine(plotter.XYs{
		{0, c}, {20, 20*x + c},
	})
	if err != nil {
		return fmt.Errorf("could not create line: %v", err)
	}
	p.Add(l)

	wt, err := p.WriterTo(256, 256, "png")
	if err != nil {
		return fmt.Errorf("could not create writer %v", err)
	}
	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to %s: %v", path, err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("could close %s: %v", path, err)
	}
	return nil
}
