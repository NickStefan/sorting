package sorting

func middleIndex(x []int)int{
	if len(x) % 2 == 0 {
		return len(x) / 2
	} else {
		return int(len(x) / 2)
	}
}

func MergeSort(x []int) []int {
	if len(x) == 1 {
		return x
	} else {
		var mid = middleIndex(x)
		var leftHalf = MergeSort(x[0:mid])
		var rightHalf = MergeSort(x[mid:len(x)])
		return merge( leftHalf, rightHalf)
	}
}

func merge(left []int, right []int) []int {
	var (
		li int
		ri int
		result = make([]int, 0, 0)
	)

	for (li < len(left) && ri < len(right)) {
		if left[li] < right[ri] {
			result = append(result, left[li])
			li++
		} else {
			result = append(result, right[ri])
			ri++
		}
	}

	return append(result, append(left[li:len(left)], right[ri:len(right)]...)... )
}