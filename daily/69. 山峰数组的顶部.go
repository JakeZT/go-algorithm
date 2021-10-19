package main

import "sort"

func main() {

}
func peakIndexInMountainArray(arr []int) int {
	res := -1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			res = i
		} else {
			break
		}
	}
	return res
}

func peakIndexInMountainArray1(arr []int) int {
	return sort.Search(len(arr)-1, func(i int) bool { return arr[i] > arr[i+1] })
}

/*

var peakIndexInMountainArray = function(arr) {
    const n = arr.length;
    let left = 1, right = n - 2, ans = 0;

    while (left <= right) {
        const mid = Math.floor((left + right) /2 );
        if (arr[mid] > arr[mid + 1]) {
            ans = mid;
            right = mid - 1;
        } else {
            left = mid + 1;
        }
    }
    return ans;
};*/
