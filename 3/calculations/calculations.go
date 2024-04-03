package calculations

import (
	"github.com/sirupsen/logrus"
)

func Calculate(n int64, log bool) int64 {
	if log {
		logrus.Println("Start calculations...")
		logrus.Printf("Calculate <%d>!\n", n)
	}

	result := int64(1)
	for i := int64(1); i <= n; i++ {
		result *= i
	}

	if log {
		logrus.Println("Calculations complete!")
	}
	return result
}
