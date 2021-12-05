package sorting

type QuickSort struct{}

func QuickSorter() QuickSort {
	return QuickSort{}
}

// pick the last num as pivot. find its right place
// repeat the process for the remaning 2 arrays
func (q QuickSort) Sort(nums []int) {

	// base case
	if len(nums) < 2 {
		return
	}

	pivotIndex := q.partition(nums, 0, len(nums)-1)

	q.Sort(nums[0:pivotIndex])
	q.Sort(nums[pivotIndex+1:])

}

// return pivot index
// bring smaller nums together at left, bigger nums at right of the pivot num
func (q QuickSort) partition(nums []int, l int, h int) int {
	pivot := nums[h] // select pivot

	sPtr := l - 1 // points to last smaller nums from pivot
	itr := l

	for itr < h {
		if nums[itr] < pivot {
			// found a smaller one, add it to smaller list
			sPtr++
			nums[sPtr], nums[itr] = nums[itr], nums[sPtr]
		}
		itr++
	}

	// put pivot into its place
	sPtr++
	nums[h], nums[sPtr] = nums[sPtr], nums[h]

	return sPtr

}
