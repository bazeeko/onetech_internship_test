package calculator

type Calculator struct {
	Input  <-chan int
	Output chan<- int
}

func (c *Calculator) Start() {
	go func(*Calculator) {
		for {
			select {
			case v, ok := <-c.Input:
				if !ok {
					close(c.Output)
					return
				}
				c.Output <- v * v
			}
		}
	}(c)
}
