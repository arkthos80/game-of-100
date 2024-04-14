package main

import (
	"fmt"
	"game-of-100/models"
	"game-of-100/strategies"
	"os"
	"strconv"
	"time"
)

func readStrategy() strategies.Strategy {
	strategyParam, err := strconv.Atoi(os.Args[1])
	if err != nil || strategyParam < 0 || strategyParam > len(strategies.All) {
		fmt.Println("invalid strategy: [", os.Args[1], "]")
		fmt.Println("Max value: ", len(strategies.All))
		os.Exit(1)
	}
	return strategies.All[strategyParam]
}

func readTimeout() time.Duration {
	timeoutParam, err := strconv.Atoi(os.Args[2])
	if err != nil || timeoutParam <= 0 {
		fmt.Println("invalid timeout: [", os.Args[2], "]")
		os.Exit(1)
	}
	return time.Duration(timeoutParam) * time.Second
}

func readStartPosition() (int, int) {
	var paramsErr []error = make([]error, 2)
	var x, y int
	x, paramsErr[0] = strconv.Atoi(os.Args[3])
	y, paramsErr[1] = strconv.Atoi(os.Args[4])
	if paramsErr[0] != nil || paramsErr[1] != nil || models.IsOutOfBound(int8(x), int8(y), 10) {
		fmt.Println("invalid starting point: [", os.Args[3], ",", os.Args[4], "]")
		os.Exit(1)
	}

	return x, y
}
