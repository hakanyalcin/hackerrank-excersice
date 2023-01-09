package main

func main() {
	ranked := []int32{100, 100, 50, 40, 40, 20, 10}
	player := []int32{5, 25, 50, 120}

	type rankBoard struct {
		score int32
		rank  int
	}
	list := []rankBoard{}
	var result []int32

	rank := 1
	for index, element := range ranked {
		list = append(list, rankBoard{rank: rank, score: element})
		if index+1 == len(ranked) {
			break
		}
		if ranked[index] != ranked[index+1] {
			rank++
		}
	}

	point := len(list) - 1

	for i := 0; i < len(player); i++ {

		for point >= 0 && player[i] > list[point].score {
			point--
		}

		if point == -1 {
			result = append(result, 1)
		} else if player[i] == list[point].score {
			result = append(result, int32(list[point].rank))
		} else if player[i] < list[point].score {
			result = append(result, int32(list[point].rank)+1)
		}

	}

	return

}
