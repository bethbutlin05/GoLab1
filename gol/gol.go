package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	newWorld := make([][]byte, p.imageHeight)
	for y := 0; y < p.imageHeight; y++ {
		newWorld[y] = make([]byte, p.imageWidth)
	}
	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			aliveNeighbours := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {

					//Don't count the cell itself
					if dx == 0 && dy == 0 {
						continue
					}

					neighbourX := (x + dx + p.imageWidth) % p.imageWidth
					neighbourY := (y + dy + p.imageHeight) % p.imageHeight

					if world[neighbourY][neighbourX] == 255 {
						aliveNeighbours++
					}
				}
			}

			if world[y][x] == 255 {
				if aliveNeighbours == 2 || aliveNeighbours == 3 {
					newWorld[y][x] = 255
				} else {
					newWorld[y][x] = 0
				}
			} else {
				if aliveNeighbours == 3 {
					newWorld[y][x] = 255
				} else {
					newWorld[y][x] = 0
				}
			}
		}
	}
	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	aliveCells := []cell{}
	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			if world[y][x] == 255 {
				aliveCells = append(aliveCells, cell{x: x, y: y})
			}
		}
	}
	return aliveCells
}
