package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	//"strings"
	"image/png"
	"io/ioutil"
	"io"
)

const ICONPATH = "images/icon.png"

func main(){
	var inTE, outTE *walk.TextEdit

	ic :=getIcon(ICONPATH)

	/*主窗体，设置标题，大小，布局和图标*/
	mainWin := MainWindow{
		Title:"DownloadManager",
		MinSize:Size{900, 480},
		Layout:  VBox{Margins:Margins{0,0,0,0}},
		Icon:ic,
	}


	/*菜单*/
	menuItems := []MenuItem{
		Menu{
			Text: "&任务",
			Items: setTaskMenu(),
		},
		Menu{
			Text: "&Help",
			Items: []MenuItem{
				Action{
					Text:        "About",
					OnTriggered: func() {
						MainWindow{
							Title:"PopWin",
							MinSize:Size{200, 180},
							Layout:  VBox{Margins:Margins{0,0,0,0}},
							//Icon:walk.NewProperty(),

						}.Run()
					},
				},
			},
		},
	}

	/*工具栏*/
	menuBody := HSplitter{
		MaxSize:Size{900, 100},
		Children: []Widget{
			TextEdit{AssignTo: &inTE},
			TextEdit{AssignTo: &outTE, ReadOnly: true},
		},
	}

	/*主内容框*/
	mainBody := HSplitter{
		MaxSize:Size{900, 380},
		Children: []Widget{
			TreeView{
				MaxSize:Size{300,380},

			},
			TextEdit{MaxSize:Size{600,380},AssignTo: &outTE, ReadOnly: true},
		},
	}

	mainWin.MenuItems = menuItems
	mainWin.Children=[]Widget{menuBody,mainBody}

	mainWin.Run();
}

func getIcon(path string) *walk.Icon {
	ioutil.ReadFile(path)
	img, err := png.Decode(io)


	ic, err := walk.NewIconFromFile(path)
	if err != nil {
		//println(err)
		return nil
	}
	return ic
}

func setTaskMenu() []MenuItem {
	return []MenuItem{
		Action{
			Text:        "添加任务",
			OnTriggered: func() {
				MainWindow{
					Title:"PopWin",
					MinSize:Size{200, 180},
					Layout:  VBox{Margins:Margins{0,0,0,0}},
					//Icon:walk.NewProperty(),

				}.Run()
			},
		},
		Action{
			Text:        "添加批量任务",
			OnTriggered: func() {
				MainWindow{
					Title:"PopWin",
					MinSize:Size{200, 180},
					Layout:  VBox{Margins:Margins{0,0,0,0}},
					//Icon:walk.NewProperty(),

				}.Run()
			},
		},
		Action{
			Text:        "从剪贴板中添加批量下载",
			OnTriggered: func() {
				MainWindow{
					Title:"PopWin",
					MinSize:Size{200, 180},
					Layout:  VBox{Margins:Margins{0,0,0,0}},
					//Icon:walk.NewProperty(),

				}.Run()
			},
		},
	}
}