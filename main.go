package packageb

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello, World!")
}

func Test() {
	logrus.Errorf("configs err: %s", "agfdfs")
}
