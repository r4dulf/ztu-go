
package main

import (
    "github.com/andlabs/ui"
    "fmt"
)

func initGUI() {
    window := ui.NewWindow("Розрахунок вартості склопакета", 400, 300, false)
    box := ui.NewVerticalBox()
    box.SetPadded(true)

    glassType := ui.NewCombobox()
    glassType.Append("Однокамерний, дерев'яний")
    glassType.Append("Двокамерний, дерев'яний")
    glassType.Append("Однокамерний, металевий")
    glassType.Append("Двокамерний, металевий")
    glassType.Append("Однокамерний, металопластиковий")
    glassType.Append("Двокамерний, металопластиковий")

    widthEntry := ui.NewEntry()
    heightEntry := ui.NewEntry()
    calcButton := ui.NewButton("Розрахувати")
    resultLabel := ui.NewLabel("")

    box.Append(ui.NewLabel("Виберіть тип склопакета:"), false)
    box.Append(glassType, false)
    box.Append(ui.NewLabel("Ширина (см):"), false)
    box.Append(widthEntry, false)
    box.Append(ui.NewLabel("Висота (см):"), false)
    box.Append(heightEntry, false)
    box.Append(calcButton, false)
    box.Append(resultLabel, false)

    window.SetChild(box)

    calcButton.OnClicked(func(*ui.Button) {
        var width, height, area, price float64
        fmt.Sscanf(widthEntry.Text(), "%f", &width)
        fmt.Sscanf(heightEntry.Text(), "%f", &height)
        area = width * height

        switch glassType.Selected() {
        case 0:
            price = area * 2.5
        case 1:
            price = area * 3.0
        case 2:
            price = area * 0.5
        case 3:
            price = area * 1.0
        case 4:
            price = area * 1.5
        case 5:
            price = area * 2.0
        }

        resultLabel.SetText(fmt.Sprintf("Вартість: %.2f грн", price))
    })

    window.OnClosing(func(*ui.Window) bool {
        ui.Quit()
        return true
    })
    window.Show()
}

func main() {
    err := ui.Main(initGUI)
    if err != nil {
        panic(err)
    }
}
