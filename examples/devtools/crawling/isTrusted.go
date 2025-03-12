package crawling

import (
	"fmt"
)

func IsTrusted(windowId int) {
	if window, ok := windows[windowId]; ok {
		page := window.energy.Page().MustWaitLoad()
		fmt.Println("TargetID:", page.TargetID)
		page.MustElement("#clickHere").MustClick()
	}
}
