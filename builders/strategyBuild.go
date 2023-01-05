package builders

var startName string
var endName string
var isAnalyzed bool

func BuildStrategy(startN, endN string, allRoutes [][]string, antsN int) ([][]string, int) {
	startName, endName, isAnalyzed = startN, endN, false
	strategies := [][][]string{}
	recursiveBuildStrategy(&strategies, [][]string{}, allRoutes)
	return selectStrategy(&strategies, antsN)
}

func selectStrategy(strategies *[][][]string, n int) ([][]string, int) {
	bestSteps, bestStrategy := -1, [][]string{}
	for _, strategy := range *strategies {
		routesAmount, s, lMax, lMin, steps, temp := len(strategy), 0, 0, len(strategy[0]), 0, 0
		for _, route := range strategy {
			if lMax < len(route) {
				lMax = len(route)
			}
			if lMin > len(route) {
				lMin = len(route)
			}
			s += len(route)
		}
		if lMax*len(strategy)-s > n {
			continue
		}
		steps += lMax - 1
		for _, route := range strategy {
			temp += lMax - len(route) + 1
		}
		steps += (n - temp) / routesAmount
		if (n-temp)%routesAmount != 0 {
			steps++
		}

		if bestSteps == -1 || bestSteps > steps {
			bestSteps = steps
			bestStrategy = strategy
			continue
		}
	}
	return bestStrategy, bestSteps
}

func recursiveBuildStrategy(strategies *[][][]string, strategy, paths [][]string) {
	if len(strategy) != 0 {
		(*strategies) = append(*strategies, strategy)
	}
	visitedNames := []string{}
	for _, path := range strategy {
		for _, room := range path {
			visitedNames = append(visitedNames, room)
		}
	}
	pathsCopy := [][]string{}
	for _, path := range paths {
		isUniq := true
		for _, room := range path {
			if isRoomVisited(room, visitedNames) {
				isUniq = false
				break
			}
		}
		if path[0] == startName && path[1] == endName {
			if isAnalyzed {
				continue
			}
			pathsCopy = append(pathsCopy, path)
			isAnalyzed = true
		}
		if isUniq {
			pathsCopy = append(pathsCopy, path)
		}
	}
	if len(pathsCopy) == 0 {
		return
	}

	for _, pathVariant := range pathsCopy {
		strategyCopy := make([][]string, len(strategy))
		copy(strategyCopy, strategy)
		strategyCopy = append(strategyCopy, pathVariant)
		recursiveBuildStrategy(strategies, strategyCopy, pathsCopy)
	}

}

func isRoomVisited(roomStr string, visitedRoomsStr []string) bool {
	for _, vR := range visitedRoomsStr {
		if vR == roomStr && vR != startName && vR != endName {
			return true
		}
	}
	return false
}
