package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"github.com/sqweek/dialog"

	"log"
	"strconv"
	"strings"

	//"github.com/ry0-suke/gofdl_tmp/"
)

func main() {
	app := app.New()
	window := app.NewWindow("File Dialog")
	window.Resize(fyne.NewSize(500, 150))

	var (
		files string //for Display
		fileNum = 0 //Num of Files
		filePath = make([]string, 0) //for Input Files
	)

	setFilePaths := widget.NewMultiLineEntry() //Entry for Selected Files
	setFilePaths.Disable()

	browseBtn := widget.NewButton("Browse", func() { //Browse Files
		file, _ := dialog.File().Title("Open Dialog").Load()
		log.Println("Browsing...")
		if strings.Index(files, string(file)) == -1 { //Check Duplicate Files
			filePath = append(filePath, string(file)) //Add File
			if files != "" {
				files += ("\n" + file)
				fileNum++
			} else {
				files = string(file)
				fileNum = 1
			}
			setFilePaths.SetText(files) //Set Files for Display
			log.Println(strconv.Itoa(fileNum) + " File Selected")
		} else {
			log.Println("Already Selected")
		}
	})

	resetBtn := widget.NewButton("Reset", func() { //Initialize Default
		log.Println("File Clear")
		files = ""
		fileNum = 0
		filePath = append(filePath[:0])
		setFilePaths.SetText(files)
	})

	execBtn := widget.NewButton("Exec", func() { //Execute Rearrange
		if fileNum != 0 { //Check Empty
			log.Println("Executing...")
			/*

			*/
			window.Close()
			log.Println("Complete")
		} else {
			log.Println("File Not Selected")
		}
	})

	window.SetContent( //Layout Contents
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			fyne.NewContainerWithLayout(layout.NewHBoxLayout(), setFilePaths, layout.NewSpacer(),fyne.NewContainerWithLayout(
					layout.NewVBoxLayout(), browseBtn, resetBtn,
				),
			),
			fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), execBtn, layout.NewSpacer(),
			),
		),
	)

	window.ShowAndRun()
}