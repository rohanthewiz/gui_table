package main

import (
	"fmt"
	"gui_table/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
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

	tbl := rtable.CreateTable(tblOpts, func(state bool, row, col int) {
		fmt.Println("Yup. Got that click! row:", row, "col:", col, "state:", state)
		rowBinding := tblOpts.Bindings[row-1]
		cellBinding, err := rowBinding.GetItem(tblOpts.ColAttrs[col].ColName)
		if err != nil {
			fmt.Println(rerr.StringFromErr(err))
			return
		}
		err = cellBinding.(binding.Bool).Set(state)
		if err != nil {
			fmt.Println(rerr.StringFromErr(err))
			return
		}
	})

	tbl.OnSelected = func(cell widget.TableCellID) {
		fmt.Println("In table handler -> Row", cell.Row, "Col", cell.Col)
		// Bounds check
		if cell.Row < 0 || cell.Row > len(data.AnimalBindings) { // 1st col is header
			fmt.Println("*-> Row out of limits")
			return
		}
		if cell.Col < 0 || cell.Col >= len(data.AnimalCols) {
			fmt.Println("*-> Column out of limits")
			return
		}

		// Handle header row clicked
		if cell.Row == 0 { // fmt.Println("-->", tblOpts.ColAttrs[cell.Col].Header)
			// Add a row
			data.AnimalBindings = append(data.AnimalBindings,
				binding.BindStruct(&data.Animal{Name: "John", Type: "Human",
					Color: "brown", Weight: "170"}))
			tblOpts.Bindings = data.AnimalBindings
			tbl.Refresh()
			return
		}

		// Handle non-header row clicked

		// The table selected event is not firing when the checkbox clicked
		// // Handle checkbox clicked
		// if cell.Col == 0 {
		// 	val, err := rtable.GetBoolCellValue(cell, tblOpts)
		// 	if err != nil {
		// 		fmt.Println(rerr.StringFromErr(err))
		// 		return
		// 	}
		// 	fmt.Println("*--> Checkbox status", val)
		// 	return
		// }

		str, err := rtable.GetStrCellValue(cell, tblOpts)
		if err != nil {
			fmt.Println(rerr.StringFromErr(err))
			return
		}

		// Get to the string binding and reverse the string
		rowBinding := tblOpts.Bindings[cell.Row-1]
		cellBinding, err := rowBinding.GetItem(tblOpts.ColAttrs[cell.Col].ColName)
		if err != nil {
			fmt.Println(rerr.StringFromErr(err))
			return
		}
		err = cellBinding.(binding.String).Set(rvsString(str))
		if err != nil {
			fmt.Println(rerr.StringFromErr(err))
			return
		}
		fmt.Println("-->", str)
	}

	// Layout
	wn.SetContent(container.NewMax(tbl))
	wn.Resize(fyne.Size{Width: 500, Height: 450})
	wn.ShowAndRun()
}

func rvsString(in string) (out string) {
	runes := []rune(in)
	ln := len(runes)
	halfLn := ln / 2

	for i := 0; i < halfLn; i++ {
		runes[i], runes[ln-1-i] = runes[ln-1-i], runes[i]
	}
	return string(runes)
}
