package aarray

//create new array with dinamic value
/*

1 2 3 => baris 0,
*/

func NewStringArray(baris int, kolom int) [][]string {
	//membuat slice, length baris
	result := make([][]string, baris)
	for i := range result {
		result[i] = make([]string, kolom)
	}
	return result
}

func NewNumberArray(baris int, kolom int) [][]int32 {
	result := make([][]int32, baris)
	for i := range result {
		result[i] = make([]int32, kolom)
	}
	return result
}
