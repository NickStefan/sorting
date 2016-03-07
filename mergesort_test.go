package sorting

import (
	"testing"
	"reflect"
	"fmt"
)

func TestMergeSort(t *testing.T){
	in := []int{
			48,96,86,68,
			57,82,63,70,
			37,34,83,27,
			19,97, 9,17,
			}
	want := []int{
			9,17,19,27,34,
			37,48,57,63,68,
			70,82,83,86,96,97,
			}

	got := MergeSort(in)
	if !reflect.DeepEqual(got, want) {
		fmt.Println("MergeSort \n", got, want)
		t.Error("MergeSort")
	}
}