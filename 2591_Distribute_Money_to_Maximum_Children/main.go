package main

func main() {
	// Run tests with "go test ."
}

const (
	PartSizeBest   = 8
	PartSizeDenied = 4
	ResultNoWay    = -1
)

func DistMoney(money int, children int) int {
	parts := money / PartSizeBest
	if parts == 0 {
		if money >= children {
			if children == 1 && money == PartSizeDenied {
				return ResultNoWay
			}
			return 0
		} else {
			return ResultNoWay
		}
	}
	if parts == children && money%PartSizeBest == 0 {
		return parts
	}
	if children > 1 {
		if parts >= children {
			parts = children - 1
			children = 1
			if parts == 0 {
				return ResultNoWay
			}
		} else {
			children = children - parts
		}
	}
	for {
		for i := 0; i < PartSizeBest; i++ {
			result := DistMoney(money-parts*PartSizeBest+i, children)
			if result >= 0 {
				if i == 0 {
					return parts
				} else {
					return parts - 1
				}
			}
		}
		parts = parts - 1
		children = children + 1
		if parts == 0 {
			return ResultNoWay
		}
	}
}
