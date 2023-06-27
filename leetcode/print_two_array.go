package leetcode

func PrintTwoArray(arr1 string, arr2 string) {
	ch := make(chan string)

	go func() {
		printFlag := true
		i := 0
		j := 0
		for i < len(arr1) || j < len(arr2) {
			if printFlag {
				if i < len(arr1) {
					ch <- string(arr1[i])
					i++
				}

				printFlag = false
			} else {
				if j < len(arr2) {
					ch <- string(arr2[j])
					j++
				}

				printFlag = true
			}
		}

		close(ch)
	}()

	for rs := range ch {
		println(rs)
	}
}
