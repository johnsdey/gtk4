package main

import (
	"os"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.example.paneclamp", 0)

	app.ConnectActivate(func() {
		win := gtk.NewApplicationWindow(app)
		win.SetTitle("Gemini Model Version 2.0")
		win.SetDefaultSize(1200, 800)
		win.SetResizable(true)
		win.SetSizeRequest(700, 400)

		mainBox := gtk.NewBox(gtk.OrientationVertical, 0)

		paned := gtk.NewPaned(gtk.OrientationHorizontal)

		// LEFT
		leftBox := gtk.NewBox(gtk.OrientationVertical, 5)
		leftLabel := gtk.NewLabel("Request")
		leftBox.Append(leftLabel)

		leftBox.SetMarginTop(10)
		leftBox.SetMarginBottom(10)
		leftBox.SetMarginStart(10)
		leftBox.SetMarginEnd(10)

		leftTextView := gtk.NewTextView()
		leftTextView.SetWrapMode(gtk.WrapWord)
		leftTextView.SetEditable(true)
		leftScrolled := gtk.NewScrolledWindow()
		leftScrolled.SetChild(leftTextView)
		leftScrolled.SetVExpand(true)
		leftScrolled.SetHExpand(true)

		leftBtnBox := gtk.NewBox(gtk.OrientationHorizontal, 5)
		for _, label := range []string{"L1", "LongLeftBtn2", "L3", "L4", "L5"} {
			leftBtnBox.Append(gtk.NewButtonWithLabel(label))
		}
		leftBtnAlign := gtk.NewBox(gtk.OrientationHorizontal, 5)
		leftBtnAlign.SetHAlign(gtk.AlignCenter)
		leftBtnAlign.Append(leftBtnBox)

		leftBox.Append(leftScrolled)
		leftBox.Append(leftBtnAlign)

		// RIGHT
		rightBox := gtk.NewBox(gtk.OrientationVertical, 5)
		rightLabel := gtk.NewLabel("Response")
		rightBox.Append(rightLabel)

		rightBox.SetMarginTop(10)
		rightBox.SetMarginBottom(10)
		rightBox.SetMarginStart(10)
		rightBox.SetMarginEnd(10)

		rightTextView := gtk.NewTextView()
		rightTextView.SetWrapMode(gtk.WrapWord)
		rightTextView.SetEditable(false)

		rightScrolled := gtk.NewScrolledWindow()
		rightScrolled.SetChild(rightTextView)
		rightScrolled.SetVExpand(true)
		rightScrolled.SetHExpand(true)

		textView := gtk.NewTextView()
		textView.SetEditable(false)
		textView.SetCursorVisible(false)
		textBuffer := textView.Buffer()
		textBuffer.SetText("This is text in the right window.\nYou can scroll, select, and even edit it.")
		rightScrolled.SetChild(textView)

		rightBtnBox := gtk.NewBox(gtk.OrientationHorizontal, 5)
		for _, label := range []string{"RBtn1", "RBtnTwoLong", "R3"} {
			rightBtnBox.Append(gtk.NewButtonWithLabel(label))
		}
		rightBtnAlign := gtk.NewBox(gtk.OrientationHorizontal, 5)
		rightBtnAlign.SetHAlign(gtk.AlignCenter)
		rightBtnAlign.Append(rightBtnBox)

		rightBox.Append(rightScrolled)
		rightBox.Append(rightBtnAlign)

		paned.SetStartChild(leftBox)
		paned.SetEndChild(rightBox)
		paned.SetPosition(400)

		statusBar := gtk.NewLabel("Status")
		statusFrame := gtk.NewFrame("")
		statusFrame.SetMarginTop(5)
		statusFrame.SetMarginBottom(5)
		statusFrame.SetMarginStart(10)
		statusFrame.SetMarginEnd(10)
		statusFrame.SetChild(statusBar)

		mainBox.Append(paned)
		mainBox.Append(statusFrame)

		win.SetChild(mainBox)

		// Connect after realize to calculate button widths

		win.SetChild(mainBox)
		win.SetVisible(true)
	})

	os.Exit(app.Run(os.Args))
}
