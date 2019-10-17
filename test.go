
package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func printStatistics() {
	elapsed := time.Duration(3333*time.Second)
	elapsedRunning := time.Duration(3000*time.Second)

	successTests := 340
	failedTests := 2
	skippedTests := 44
	timeoutTests := 3

	failedNames := "\n\t\ttest1\n\t\ttest2"

	logrus.Infof("Statistics:" +
		fmt.Sprintf("\n\tElapsed total: %v", elapsed) +
		fmt.Sprintf("\n\tTests time: %v", elapsedRunning) +
		fmt.Sprintf("\n\tTasks  Completed: %d", 445) +
		fmt.Sprintf("\n\tStatus  Passed: %d"+
			"\n\tStatus  Failed: %d%v"+
			"\n\tStatus  Timeout: %d"+
			"\n\tStatus  Skipped: %d", successTests, failedTests, failedNames, timeoutTests, skippedTests))
}

func main() {
	printStatistics()
}