package roundrobin

import (
	"sync"
)

type RoundData struct {
	Data   interface{}
	Weight int
}

type RoundRobin struct {
	data       []RoundData
	lastHit    int // 上次选中的index
	currWeight int // 当前选中的权重值
	gcd        int //当前所有权重的最大公约数 比如 2，4，8 的最大公约数为：2
	maxWeight  int // 最大权重值

	lock *sync.RWMutex
}

func NewRoundRobin(data []RoundData) *RoundRobin {
	r := &RoundRobin{}
	r.lock = new(sync.RWMutex)
	r.data = data
	r.gcd = r.getGcd()
	r.maxWeight = r.getMaxWeight()
	r.lastHit = -1

	return r
}

// 取值
func (this *RoundRobin) Get() interface{} {

	this.lock.RLock()

	defer this.lock.RUnlock()
	n := len(this.data)

	if n == 1 {
		return this.data[0].Data
	}

	for {
		this.lastHit = (this.lastHit + 1) % n
		if this.lastHit == 0 {
			this.currWeight = this.currWeight - this.gcd
			if this.currWeight <= 0 {
				this.currWeight = this.maxWeight
			}

		}

		if this.data[this.lastHit].Weight >= this.currWeight {
			return this.data[this.lastHit].Data
		}
	}

	// 不会执行到这
	return nil
}

// 获取最大的权值
func (this *RoundRobin) getMaxWeight() int {
	max := 0
	for i, _ := range this.data {
		if this.data[i].Weight > max {
			max = this.data[i].Weight
		}
	}

	return max
}

//  获取服务器所有权值的最大公约数
func (this *RoundRobin) getGcd() int {
	ints := make([]int, len(this.data))
	for i, _ := range this.data {
		ints[i] = this.data[i].Weight
	}

	return ngcd(ints)
}

func (this *RoundRobin) Reset(data []RoundData) {
	this.lock.Lock()

	this.data = data
	this.gcd = this.getGcd()
	this.maxWeight = this.getMaxWeight()
	this.lastHit = -1

	this.lock.Unlock()
}
