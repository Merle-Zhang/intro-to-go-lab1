package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	tmp := make([][]byte, len(world))
	for i := range world {
		tmp[i] = make([]byte, len(world[i]))
		copy(tmp[i], world[i])
	}

	for i := range world {
		for j := range world[i] {
			count := countLiveNeighbours(i, j, world)
			if world[i][j] == 255 && (count < 2 || count > 3) {
				tmp[i][j] = 0
			} else if world[i][j] == 0 && count == 3 {
				tmp[i][j] = 255
			}
		}
	}

	return tmp
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	cells := []cell{}
	for i := range world {
		for j := range world[i] {
			if world[i][j] == 255 {
				cells = append(cells, cell{x: j, y: i})
			}
		}
	}
	return cells
}

func countLiveNeighbours(i int, j int, world [][]byte) int {
	lst := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	count := 0
	for _, r := range lst {
		if world[(i+r[0]+16)%16][(j+r[1]+16)%16] == 255 {
			count++
		}
	}
	return count
}
