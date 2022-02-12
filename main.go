package main

import (
	"fmt"
	"gui_table/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/rohanthewiz/rerr"
	"github.com/rohanthewiz/rtable"
)

func main() {
	ap := app.New()
	wn := ap.NewWindow("Table Widget")

	tblOpts := &rtable.TableOptions{
		RefWidth: "reference width",
		ColAttrs: data.AnimalCols,
		Bindings: data.AnimalBindings,
	}

	tbl := rtable.CreateTable(tblOpts,
		func(cell widget.TableCellID) {
			if cell.Row == 0 && cell.Col >= 0 && cell.Col < len(data.AnimalCols) { // valid hdr cell
				fmt.Println("-->", tblOpts.ColAttrs[cell.Col].Header)
				return
			}

			str, err := rtable.GetStrCellValue(cell, tblOpts)
			if err != nil {
				fmt.Println(rerr.StringFromErr(err))
				return
			}
			fmt.Println("-->", str)
		})

	// Layout
	wn.SetContent(container.NewMax(tbl))
	wn.Resize(fyne.Size{Width: 500, Height: 450})
	wn.ShowAndRun()
}
