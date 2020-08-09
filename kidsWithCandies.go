package kis_with_candies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	lenOfCandies := len(candies)
	ret := make([]bool, lenOfCandies)
	max := 0
	for _, val := range candies {
		if val > max {
			max = val
		}
	}
	for idx, val := range candies {
		ret[idx] = (val + extraCandies) >= max
	}
	return ret
}
