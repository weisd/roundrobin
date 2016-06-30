package roundrobin

import (
	"sync"
)

type RoundData struct {
	Data   interface{}
	Weight int
}

type RoundRobin struct {
	Data       []RoundData
	LastHit    int // 上次选中的index
	CurrWeight int // 当前选中的权重值
	Gcd        int //当前所有权重的最大公约数 比如 2，4，8 的最大公约数为：2
	MaxWeight  int // 最大权重值

	lock *sync.RWMutex
}

func NewRoundRobin(data []RoundData) *RoundRobin {
	r := &RoundRobin{}
	r.lock = new(sync.RWMutex)
	r.Data = data
	r.Gcd = r.getGcd()
	r.MaxWeight = r.getMaxWeight()
	r.LastHit = -1

	return r
}

// 取值
func (this *RoundRobin) Get() interface{} {
	this.lock.RLock()

	defer this.lock.RUnlock()

	n := len(this.Data)
	for {
		this.LastHit = (this.LastHit + 1) % n
		if this.LastHit == 0 {
			this.CurrWeight = this.CurrWeight - this.Gcd
			if this.CurrWeight <= 0 {
				this.CurrWeight = this.MaxWeight
			}

		}

		if this.Data[this.LastHit].Weight >= this.CurrWeight {
			return this.Data[this.LastHit]
		}
	}

	// 不会执行到这
	return nil
}

// 获取最大的权值
func (this *RoundRobin) getMaxWeight() int {
	max := 0
	for i, _ := range this.Data {
		if this.Data[i].Weight > max {
			max = this.Data[i].Weight
		}
	}

	return max
}

//  获取服务器所有权值的最大公约数
func (this *RoundRobin) getGcd() int {
	ints := make([]int, len(this.Data))
	for i, _ := range this.Data {
		ints[i] = this.Data[i].Weight
	}

	return ngcd(ints)
}

func (this *RoundRobin) Reset(data []RoundData) {
	this.lock.Lock()

	this.Data = data
	this.Gcd = this.getGcd()
	this.MaxWeight = this.getMaxWeight()
	this.LastHit = -1

	this.lock.Unlock()
}
