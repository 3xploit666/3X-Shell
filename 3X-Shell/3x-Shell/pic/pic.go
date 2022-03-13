package pic

import (
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
)

func GetScreenshot() {
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := "c:\\users\\public\\cap.png"
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)
	}
}
