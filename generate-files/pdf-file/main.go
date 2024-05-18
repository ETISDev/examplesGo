
package main
import (
	"log"
	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4 })
	pdf.AddPage()
	var err error
	err = pdf.AddTTFFont("loma", "./res/LiberationSerif-Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.Image("./res/gopher01.jpg", 200, 50, nil) //print image
	err = pdf.SetFont("loma", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetXY(250, 200) //move current location
	pdf.Cell(nil, "gopher and gopher") //print text

	pdf.WritePdf("image.pdf")
}