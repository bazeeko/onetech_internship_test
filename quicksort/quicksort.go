package quicksort

func quicksort(array []int, low int, high int) {
	var l int = low
	var h int = high
	var pivot = array[(l+h)/2]

	for l <= h {
		for array[l] < pivot {
			l++
		}
		for array[h] > pivot {
			h--
		}

		if l <= h {
			array[l], array[h] = array[h], array[l]
			l++
			h--
		}
	}
	if h > low {
		quicksort(array, low, h)
	}

	if l < high {
		quicksort(array, l, high)
	}
}

func QuickSort(a []int) {
	quicksort(a, 0, len(a)-1)
}
