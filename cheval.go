package main

/*
i5 2400 177ms
i3 530 350ms
*/

import (
	"fmt"
	"sync"
	"time"
)

const nbmove = 8

const SIDE = 5
const SQR_SIDE = SIDE * SIDE

var movex = []int{-1, -1, -2, -2, +1, +1, +2, +2}
var movey = []int{-2, +2, -1, +1, +2, -2, +1, -1}
var shift = [8]int{}
var shift_0 int
var shift_1 int
var shift_2 int
var shift_3 int
var shift_4 int
var shift_5 int
var shift_6 int
var shift_7 int

var nbcoup = 1
var nbsol = 0

func main() {
	t := time.Now()
	start()
	fmt.Println(time.Since(t))
}
func start() {
	tstart := time.Now()
	for i := 0; i < 8; i++ {
		shift[i] = movex[i]*SIDE + movey[i]
	}
	shift_0 = shift[0]
	shift_1 = shift[1]
	shift_2 = shift[2]
	shift_3 = shift[3]
	shift_4 = shift[4]
	shift_5 = shift[5]
	shift_6 = shift[6]
	shift_7 = shift[7]
	fmt.Println(shift)
	fmt.Printf("%dx%d\n", SIDE, SIDE)
	var wg sync.WaitGroup
	for xs := 0; xs < SIDE; xs++ {
		for ys := 0; ys < SIDE; ys++ {
			wg.Add(1)
			circuit := make([]int, SQR_SIDE)
			go do_solve(&wg, circuit, 1, xs, ys)
		}
	}
	//circuit := make([]int,SQR_SIDE)
	//solve(&wg,circuit,1,0,0)
	wg.Wait()
	duration := time.Since(tstart)
	fmt.Printf("%s nbsol=%d\n", duration, nbsol)
}

func do_solve(wg *sync.WaitGroup, circuit []int, nb int, x int, y int) {
	solve(circuit, nb, x, y)
	wg.Done()
	return
}
func solve(circuit []int, nb int, x int, y int) {
	//paint()
	var newx int
	var newy int
	//fmt.Println(nb, x,y)
	pos := x*SIDE + y
	circuit[pos] = nb
	//paint (circuit)
	if nb == SQR_SIDE {
		//paint(circuit)
		nbsol += 1
		circuit[pos] = 0
		return
	}
	nb++
	newx = x - 1
	if newx >= 0 && newx < SIDE {
		newy = y - 2
		if newy >= 0 && newy < SIDE && circuit[pos+shift_0] == 0 {
			solve(circuit, nb, newx, newy)
		}
		newy = y + 2
		if newy >= 0 && newy < SIDE && circuit[pos+shift_1] == 0 {
			solve(circuit, nb, newx, newy)
		}
	}

	newx = x - 2
	if newx >= 0 && newx < SIDE {
		newy = y - 1
		if newy >= 0 && newy < SIDE && circuit[pos+shift_2] == 0 {
			solve(circuit, nb, newx, newy)
		}
		newy = y + 1
		if newy >= 0 && newy < SIDE && circuit[pos+shift_3] == 0 {
			solve(circuit, nb, newx, newy)
		}
	}
	newx = x + 1
	if newx >= 0 && newx < SIDE {
		newy = y + 2
		if newy >= 0 && newy < SIDE && circuit[pos+shift_4] == 0 {
			solve(circuit, nb, newx, newy)
		}
		newy = y - 2
		if newy >= 0 && newy < SIDE && circuit[pos+shift_5] == 0 {
			solve(circuit, nb, newx, newy)
		}
	}
	newx = x + 2
	if newx >= 0 && newx < SIDE {
		newy = y + 1
		if newy >= 0 && newy < SIDE && circuit[pos+shift_6] == 0 {
			solve(circuit, nb, newx, newy)
		}
		newy = y - 1
		if newy >= 0 && newy < SIDE && circuit[pos+shift_7] == 0 {
			solve(circuit, nb, newx, newy)
		}
	}

	circuit[pos] = 0
	return
}

func paint(circuit []int) {
	fmt.Println(nbsol)
	for x := 0; x < SIDE; x++ {
		for y := 0; y < SIDE; y++ {
			//   fmt.Println(x,y)
			//  fmt.Println(x*SIDE+y)
			if SQR_SIDE < 100 {
				fmt.Printf("%02d ", circuit[x*SIDE+y])
			} else {
				fmt.Printf("%03d ", circuit[x*SIDE+y])
			}
		}
		fmt.Println("")
	}
}
