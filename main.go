package main

// Code by Rohan Allison
import (
	"fmt"
	"gui_table/table"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/rohanthewiz/rerr"
)

type Animal struct {
	Name, Type, Color, Weight string
}

var animals = []Animal{
	{Name: "Frisky", Type: "cat", Color: "gray", Weight: "10"},
	{Name: "Ella", Type: "dog", Color: "brown", Weight: "50"},
	{Name: "Mickey", Type: "mouse", Color: "black", Weight: "1"},
}

var animalCols = []table.ColAttr{
	{ColName: "Name", Header: "Name", WidthPercent: 100},
	{ColName: "Type", Header: "Type", WidthPercent: 67},
	{ColName: "Color", Header: "Color", WidthPercent: 100},
	{ColName: "Weight", Header: "Weight", WidthPercent: 67},
}

var animalBindings []binding.DataMap

// Create a binding for each animal data
func init() {
	for i := 0; i < len(animals); i++ {
		animalBindings = append(animalBindings, binding.BindStruct(&animals[i]))
	}
}

func main() {
	ap := app.New()
	wn := ap.NewWindow("Table Widget")

	tblOpts := &table.TableOptions{
		RefWidth: "reference width",
		ColAttrs: animalCols,
		Bindings: animalBindings,
	}

	tbl := table.CreateTable(tblOpts,
		func(cell widget.TableCellID) {
			if cell.Row == 0 && cell.Col >= 0 && cell.Col < len(animalCols) { // valid hdr cell
				fmt.Println("-->", tblOpts.ColAttrs[cell.Col].Header)
				return
			}

			str, err := table.GetStrCellValue(cell, tblOpts)
			if err != nil {
				fmt.Println(rerr.StringFromErr(err))
				return
			}
			fmt.Println("-->", str)
		})

	// Layout
	refWidth := widget.NewLabel(tblOpts.RefWidth).MinSize().Width

	for i := 0; i < len(animalCols); i++ {
		tbl.SetColumnWidth(i, float32(animalCols[i].WidthPercent)/100.0*refWidth)
	}

	wn.SetContent(container.NewMax(tbl))
	wn.Resize(fyne.Size{Width: 500, Height: 450})
	wn.ShowAndRun()
}
