package crawling

import (
	"fmt"
	"github.com/go-rod/rod/lib/proto"
	"github.com/ysmood/gson"
	"io/ioutil"
)

func Screenshot(windowId int) {
	if window, ok := windows[windowId]; ok {
		page := window.energy.Page().MustWaitLoad()
		fmt.Println("TargetID:", page.TargetID)

		buf, err := page.Screenshot(true, &proto.PageCaptureScreenshot{
			Format:  proto.PageCaptureScreenshotFormatJpeg,
			Quality: gson.Int(90),
		})
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile("fullScreenshot.png", buf, 0o644)
		if err != nil {
			panic(err)
		}
	}
}
