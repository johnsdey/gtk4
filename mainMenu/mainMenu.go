package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

var currentFilename string

func main() {
	app := gtk.NewApplication("com.example.gtkjson", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() {
		buildUI(app)
	})
	app.Run(os.Args)
}

func buildUI(app *gtk.Application) {
	win := gtk.NewApplicationWindow(app)
	win.SetTitle("GTK4 JSON File Menu")
	win.SetDefaultSize(800, 600)

	box := gtk.NewBox(gtk.OrientationVertical, 0)
	win.SetChild(box)

	// --- Menu Buttons ---
	fileMenu := gio.NewMenu()
	fileMenu.Append("New", "app.new")
	fileMenu.Append("Open", "app.open")
	fileMenu.Append("Save", "app.save")
	fileMenu.Append("Save As", "app.saveas")

	fileButton := gtk.NewMenuButton()
	fileButton.SetLabel("File")
	fileButton.SetMenuModel(fileMenu)

	viewMenu := gio.NewMenu()
	viewMenu.Append("Show Status", "app.status")

	viewButton := gtk.NewMenuButton()
	viewButton.SetLabel("View")
	viewButton.SetMenuModel(viewMenu)

	menuBar := gtk.NewBox(gtk.OrientationHorizontal, 6)
	menuBar.Append(fileButton)
	menuBar.Append(viewButton)
	box.Append(menuBar)

	// --- Actions ---
	// New
	newAction := gio.NewSimpleAction("new", nil)
	newAction.ConnectActivate(func(param *glib.Variant) {
		currentFilename = ""
		fmt.Println("New file")
	})
	app.AddAction(newAction)

	// Open
	openAction := gio.NewSimpleAction("open", nil)
	openAction.ConnectActivate(func(param *glib.Variant) {
		file, err := runZenityFileChooser("open", "Open JSON File")
		if err != nil {
			fmt.Println("Open canceled or error:", err)
			return
		}
		currentFilename = file
		fmt.Println("Opened:", file)
	})
	app.AddAction(openAction)

	// Save
	saveAction := gio.NewSimpleAction("save", nil)
	saveAction.ConnectActivate(func(param *glib.Variant) {
		if currentFilename == "" {
			fmt.Println("No filename set. Use Save As.")
			return
		}
		fmt.Println("Saved to:", currentFilename)
		// Add file writing logic here
	})
	app.AddAction(saveAction)

	// Save As
	saveAsAction := gio.NewSimpleAction("saveas", nil)
	saveAsAction.ConnectActivate(func(param *glib.Variant) {
		file, err := runZenityFileChooser("save", "Save JSON File")
		if err != nil {
			fmt.Println("Save As canceled or error:", err)
			return
		}
		if !strings.HasSuffix(file, ".json") {
			file += ".json"
		}
		currentFilename = file
		fmt.Println("Saved As:", file)
		// Add file writing logic here
	})
	app.AddAction(saveAsAction)

	// Status
	statusAction := gio.NewSimpleAction("status", nil)
	statusAction.ConnectActivate(func(param *glib.Variant) {
		fmt.Println("Status requested.")
	})
	app.AddAction(statusAction)

	win.Show()
}

func runZenityFileChooser(mode string, title string) (string, error) {
	var args []string
	args = append(args, "--file-selection", "--title="+title)
	if mode == "save" {
		args = append(args, "--save", "--confirm-overwrite")
	}
	// Limit to JSON files
	args = append(args, "--file-filter=JSON files ( *.json ) | *.json")

	cmd := exec.Command("zenity", args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	// Remove trailing newline
	return strings.TrimSpace(string(output)), nil
}
