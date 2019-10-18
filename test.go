
package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	msg := "Statistics:" +
		"\n\tStatus  Passed: 7" +
		"\n\tStatus  Failed: 3" +
		"\n\t\ttest1\n\t\ttest2"

	//fmt.Println("=================== via fmt.Println:")
	//fmt.Println(msg)

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
	})

	fmt.Println("=================== via logrus.Infof:")
	logrus.Info(msg)
}
