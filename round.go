package roundrobin

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

}

func NewRoundRobin(data []RoundData) *RoundRobin {
	r := &RoundRobin{}
	r.Data = data
	r.Gcd = r.GetGcd()
	r.MaxWeight = r.GetMaxWeight()
	r.LastHit = -1

	return r
}

// 最值
func (this *RoundRobin) GetServer() interface{} {
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
}

// 获取最大的权值
func (this *RoundRobin) GetMaxWeight() int {
	max := 0
	for i, _ := range this.Data {
		if this.Data[i].Weight > max {
			max = this.Data[i].Weight
		}
	}

	return max
}

//  获取服务器所有权值的最大公约数
func (this *RoundRobin) GetGcd() int {
	ints := make([]int, len(this.Data))
	for i, _ := range this.Data {
		ints[i] = this.Data[i].Weight
	}

	return ngcd(ints)
}

// 两个数的最大公约数 欧几里德算法
func gcd(a, b int) int {

	if a < b {
		a, b = b, a
	}

	if b == 0 {
		return a

	}
	return gcd(b, a%b)
}

// n个数的最大公约数算法
// 说明:
// 把n个数保存为一个数组
// 参数为数组的指针和数组的大小(需要计算的数的个数)
// 然后先求出gcd(a[0],a[1]), 然后将所求的gcd与数组的下一个元素作为gcd的参数继续求gcd
// 这样就产生一个递归的求ngcd的算法
func ngcd(ints []int) int {

	n := len(ints)

	if n == 1 {
		return ints[0]
	}

	return gcd(ints[n-1], ngcd(ints[0:n-1]))
}
