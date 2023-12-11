package first

func Mergesort(arry []int) []int {
	//lõpptingimus recursivil, kui
	if len(arry) < 2 {
		return arry
	}
	//teed array kaheks
	leftside := arry[:len(arry)/2]
	// ei ole len -1 kuna array puhul loetakse kuni viimaseni
	rightside := arry[len(arry)/2:]

	// kutsudes seda merge sorti jälle esile, jagame me arrayd seni kaua pooleks kuni saame array
	//pikkusega 1 ja siis Mergesort retunib selle meile merge functionile, mis siis võrdleb ja liidab
	//array lõpuks jälle kokku.
	return Merge(Mergesort(leftside), Mergesort(rightside))

}

// hakkab võrdlema juba sorditud arraysid alustades arraydest pikkusega 1, mis on juba sorditud
// kuna need on ainukesed
func Merge(left, right []int) []int {
	resultArray := []int{}
	leftIndex := 0
	rightIndex := 0

	// senikaua jookseb kuni kas parem või vasaku poole indeksi lugeja muutub suuremaks kui on array pikkus
	// kui mõlema array pikkus on 1 siis võrdleb ühe korra ja lisab väiksema resultArraysse.
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			resultArray = append(resultArray, left[leftIndex])
			leftIndex++
		} else {
			resultArray = append(resultArray, right[rightIndex])
			rightIndex++
		}
		// lisab need elemedid resltArraysse mis üle jäid (üks pool lõpeb alati varem ära)
	}
	for ; leftIndex < len(left); leftIndex++ {
		resultArray = append(resultArray, left[leftIndex])
	}
	for ; rightIndex < len(right); rightIndex++ {
		resultArray = append(resultArray, right[rightIndex])
	}
	return resultArray
}
