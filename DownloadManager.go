package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	//"strings"
)

func main(){
	var inTE, outTE *walk.TextEdit

	/*主窗体，设置标题，大小，布局和图标*/
	mainWin := MainWindow{
		Title:"DownloadManager",
		MinSize:Size{900, 480},
		Layout:  VBox{Margins:Margins{0,0,0,0}},
		Icon:getIcon(),
	}


	/*菜单*/
	menuItems := []MenuItem{
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
		MaxSize:Size{900, 120},
		Children: []Widget{
			TextEdit{AssignTo: &inTE},
			TextEdit{AssignTo: &outTE, ReadOnly: true},
		},
	}

	/*主内容框*/
	mainBody := HSplitter{
		MaxSize:Size{900, 360},
		Children: []Widget{
			PushButton{
				OnClicked:func(){
					MainWindow{
						Title:"PopWin",
						MinSize:Size{200, 180},
						Layout:  VBox{Margins:Margins{0,0,0,0}},
						//Icon:walk.NewProperty(),

					}.Run()
				},
			},
		},
	}

	mainWin.MenuItems = menuItems
	mainWin.Children=[]Widget{menuBody,mainBody}

	mainWin.Run();
}

func getIcon(path string) *walk.Icon {

	return nil
}