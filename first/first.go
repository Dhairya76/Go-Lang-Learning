package first

import (
	"fmt"
	"time"
)

func Hello() {
	fmt.Println("The conference is scheduled at: ", time.Now().Format("2006-January-02"))
}
