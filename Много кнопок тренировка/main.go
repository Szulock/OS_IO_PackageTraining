package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image/color"
	"os"
)

var task1 widget.Clickable
var task1Writer widget.Editor
var task2Button widget.Clickable
var task3Writer widget.Editor
var task4Writer widget.Editor

type AppState struct {
	currentInterface func(th *material.Theme, gtx C) D
	// Добавляем состояние для хранения карт

}
type C = layout.Context
type D = layout.Dimensions

func main() {
	state := &AppState{}
	MainWindow := new(app.Window)
	MainWindow.Option(
		app.Size(unit.Dp(1000), unit.Dp(800)),
		app.Title("BlackJack Super"),
		app.MinSize(unit.Dp(1000), unit.Dp(800)),
		// Установка максимального размера окна
		app.MaxSize(unit.Dp(1000), unit.Dp(800)),
	)

	state.currentInterface = SeeMainWindow
	RunMainWindow(MainWindow, state)
}

func RunMainWindow(window *app.Window, state *AppState) {

	th := material.NewTheme()
	var ops op.Ops
	for {

		e := window.Event()
		switch e := e.(type) {
		case app.DestroyEvent:
			return
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			// Отрисовка текущего интерфейса

			state.currentInterface(th, gtx)

			e.Frame(gtx.Ops)

		}
	}
}

func funcionalitiestaskone(gtx C) {
	var fileCreate *os.File
	var err error

	// Task1Funcionality:
	if task1.Clicked(gtx) {
		fileCreate, err = os.Create("Example.txt")
		if err != nil {
			fmt.Println("Не получилось создать файл:", err)
			return
		}
		defer fileCreate.Close() // Убедитесь, что файл будет закрыт при выходе из функции
	}

	for {
		e, ok := gtx.Event(
			key.Filter{
				Focus: &task1Writer,
				Name:  key.NameReturn,
			},
		)
		if !ok {
			break
		}
		ev, ok := e.(key.Event)
		if !ok {
			continue
		}
		// Обработка нажатия Enter
		if ev.State == key.Press && ev.Name == key.NameReturn {
			textToWrite := task1Writer.Text() // Получаем текст из редактора
			err := writeToFile("Example.txt", textToWrite)
			if err != nil {
				fmt.Println("Ошибка при записи в файл:", err)
				return
			}
		}
	}
}
func funcionalitiestaskTwo(gtx C) {
	if task2Button.Clicked(gtx) {
		Copyfile("Example.txt", "FreshExample.txt")
		fmt.Println("Успешно создано: FreshExample.txt")
	}
}

func funcionalitiestaskThree(gtx C) {
	for {
		e, ok := gtx.Event(
			key.Filter{
				Focus: &task3Writer,
				Name:  key.NameReturn,
			},
		)
		if !ok {
			break
		}
		ev, ok := e.(key.Event)
		if !ok {
			continue
		}
		// Обработка нажатия Enter
		if ev.State == key.Press && ev.Name == key.NameReturn {
			Directory := task3Writer.Text() // Получаем текст из редактора
			IsExist(Directory)
		}
	}
}

func funcionalitiestaskfourth(gtx C) {
	for {
		e, ok := gtx.Event(
			key.Filter{
				Focus: &task4Writer,
				Name:  key.NameReturn,
			},
		)
		if !ok {
			break
		}
		ev, ok := e.(key.Event)
		if !ok {
			continue
		}
		// Обработка нажатия Enter
		if ev.State == key.Press && ev.Name == key.NameReturn {
			Directory := task4Writer.Text() // Получаем текст из редактора
			allFiles(Directory)

		}
	}
}

func writeToFile(filename, text string) error {
	fileOpened, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fileOpened.Close() // Убедитесь, что файл будет закрыт при выходе из функции

	_, err = fileOpened.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

func SeeMainWindow(th *material.Theme, gtx C) D {
	funcionalitiestaskfourth(gtx)
	funcionalitiestaskone(gtx)
	funcionalitiestaskThree(gtx)
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// Первый элемент: текст и кнопка для создания файла
		layout.Rigid(func(gtx C) D {
			return MakeTextAboveAndButton(th, gtx, "Создание файла", unit.Sp(11), &task1, "Создать файл Example.txt")
		}),
		layout.Rigid(func(gtx C) D {
			insetOut := layout.Inset{Top: unit.Dp(20)}
			return insetOut.Layout(gtx, func(gtx C) D {
				border := widget.Border{
					Color:        color.NRGBA{R: 155, G: 55, B: 155, A: 255},
					CornerRadius: unit.Dp(5),
					Width:        unit.Dp(4),
				}
				return border.Layout(gtx, func(gtx C) D {
					inset := layout.Inset{
						Top:    unit.Dp(10),
						Right:  unit.Dp(10),
						Bottom: unit.Dp(10),
						Left:   unit.Dp(10),
					}
					return inset.Layout(gtx, func(gtx C) D {
						userInput := material.Editor(th, &task1Writer, "Введите содержимое файла")
						return userInput.Layout(gtx)
					})
				})
			})
		}),
		layout.Rigid(func(gtx C) D {
			funcionalitiestaskTwo(gtx)
			return MakeTextAboveAndButton(th, gtx, "Копирование файла", unit.Sp(11), &task2Button, "Копировать файл Example и создать новый")
		}),
		layout.Rigid(func(gtx C) D {
			lbl := material.Body1(th, "Проверка на наличие файла/директории: введите путь к файлу")
			lbl.TextSize = unit.Sp(14)
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx C) D {
			insetOut := layout.Inset{Top: unit.Dp(20)}
			return insetOut.Layout(gtx, func(gtx C) D {
				border := widget.Border{
					Color:        color.NRGBA{R: 155, G: 55, B: 155, A: 255},
					CornerRadius: unit.Dp(5),
					Width:        unit.Dp(4),
				}
				return border.Layout(gtx, func(gtx C) D {
					inset := layout.Inset{
						Top:    unit.Dp(10),
						Right:  unit.Dp(10),
						Bottom: unit.Dp(10),
						Left:   unit.Dp(10),
					}
					return inset.Layout(gtx, func(gtx C) D {
						userInput3 := material.Editor(th, &task3Writer, "Введите адрес файла")
						return userInput3.Layout(gtx)
					})
				})
			})
		}),
		layout.Rigid(func(gtx C) D {
			StandardInset := layout.Inset{
				Top: unit.Dp(22),
			}
			return StandardInset.Layout(gtx, func(gtx C) D {
				lbl := material.Body1(th, "Перечисление всех файлов в директории")
				lbl.TextSize = unit.Sp(14)
				return lbl.Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx C) D {
			insetOut := layout.Inset{Top: unit.Dp(3)}
			return insetOut.Layout(gtx, func(gtx C) D {
				border := widget.Border{
					Color:        color.NRGBA{R: 155, G: 55, B: 155, A: 255},
					CornerRadius: unit.Dp(5),
					Width:        unit.Dp(4),
				}
				return border.Layout(gtx, func(gtx C) D {
					inset := layout.Inset{
						Top:    unit.Dp(10),
						Right:  unit.Dp(10),
						Bottom: unit.Dp(10),
						Left:   unit.Dp(10),
					}
					return inset.Layout(gtx, func(gtx C) D {
						userInput4 := material.Editor(th, &task4Writer, "Введите адрес директории")
						return userInput4.Layout(gtx)
					})
				})
			})
		}),
	)
}

func MakeTextAboveAndButton(th *material.Theme, gtx C, textAbove string, size unit.Sp, button *widget.Clickable, textInBtn string) D {

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			lbl := material.Body1(th, textAbove)
			lbl.TextSize = size
			return lbl.Layout(gtx)
		}),
		layout.Rigid(func(gtx C) D {
			btn := material.Button(th, button, textInBtn)
			btn.Background = color.NRGBA{R: 189, G: 25, B: 19, A: 255}
			btn.CornerRadius = unit.Dp(1)

			border := widget.Border{
				Color:        color.NRGBA{R: 255, G: 255, B: 255, A: 255},
				CornerRadius: unit.Dp(5),
				Width:        unit.Dp(7),
			}

			insets := layout.Inset{
				Top:    border.Width,
				Right:  border.Width,
				Bottom: border.Width,
				Left:   border.Width,
			}

			return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return insets.Layout(gtx, btn.Layout)
			})
		}),
	)
}
