
package main

import (
    "github.com/andlabs/ui"
    "fmt"
)

func initGUI() {
    window := ui.NewWindow("Розрахунок вартості туру", 400, 300, false)
    box := ui.NewVerticalBox()
    box.SetPadded(true)

    country := ui.NewCombobox()
    country.Append("Болгарія, літо")
    country.Append("Болгарія, зима")
    country.Append("Німеччина, літо")
    country.Append("Німеччина, зима")
    country.Append("Польща, літо")
    country.Append("Польща, зима")

    daysEntry := ui.NewEntry()
    guideCheckbox := ui.NewCheckbox("Індивідуальний гід")
    luxuryCheckbox := ui.NewCheckbox("Проживання у люксі")
    calcButton := ui.NewButton("Розрахувати")
    resultLabel := ui.NewLabel("")

    box.Append(ui.NewLabel("Виберіть країну та сезон:"), false)
    box.Append(country, false)
    box.Append(ui.NewLabel("Кількість днів:"), false)
    box.Append(daysEntry, false)
    box.Append(guideCheckbox, false)
    box.Append(luxuryCheckbox, false)
    box.Append(calcButton, false)
    box.Append(resultLabel, false)

    window.SetChild(box)

    calcButton.OnClicked(func(*ui.Button) {
        var days int
        var total float64
        fmt.Sscanf(daysEntry.Text(), "%d", &days)

        pricePerDay := 0
        switch country.Selected() {
        case 0:
            pricePerDay = 100
        case 1:
            pricePerDay = 150
        case 2:
            pricePerDay = 160
        case 3:
            pricePerDay = 200
        case 4:
            pricePerDay = 120
        case 5:
            pricePerDay = 180
        }

        total = float64(days * pricePerDay)
        if guideCheckbox.Checked() {
            total += float64(days * 50)
        }
        if luxuryCheckbox.Checked() {
            total *= 1.2
        }

        resultLabel.SetText(fmt.Sprintf("Вартість туру: $%.2f", total))
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
