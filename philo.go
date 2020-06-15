package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	id              int
	permit			chan bool
}

func (p Philo) eat(wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {

		p.permit <- true
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat ", p.id)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("finishing eating ", p.id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		<- p.permit
	}

	defer wg.Done()
}

func main() {

	Csticks := make([]*ChopS, 5)
	var wg sync.WaitGroup
	permit := make(chan bool, 2)

	for i := 0; i < 5; i++ {
		Csticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{Csticks[i], Csticks[(i+1)%5], i + 1, permit}

	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(&wg)
	}
	wg.Wait()

	fmt.Println("DONE")

}
