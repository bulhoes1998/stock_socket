// Demo code for the Table primitive.
package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(true)
	lorem := strings.Split("Ativo Ultimo Data/Hora Variacao Maximo Minimo Abertura Fechamento Of.Compra Of.Venda Media", " ")
	cols := 10
	word := 0

	for c := 0; c < cols; c++ {
		color := tcell.ColorWhite
		table.SetCell(0, c,
			tview.NewTableCell(lorem[word]).
				SetTextColor(color).
				SetAlign(tview.AlignCenter))
		word = (word + 1) % len(lorem)
	}
	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
	})
	if err := app.SetRoot(table, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
