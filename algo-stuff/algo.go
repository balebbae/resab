package main

import (
    "fmt"

    "github.com/balebbae/resaB/jsonmanager"
    "github.com/balebbae/resaB/scheduler"
)

func main() {
    // Initialize JsonManager
    jm := jsonmanager.New("input.json", "output.json")

    // Generate the schedule
    schedule, err := scheduler.ScheduleShifts(jm)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the schedule
    fmt.Println("Final Schedule:")
    for day, positions := range schedule {
        fmt.Printf("%s:\n", day)
        for position, shifts := range positions {
            fmt.Printf("  %s:\n", position)
            for shift, employee := range shifts {
                fmt.Printf("    Shift %s: %s\n", shift, employee)
            }
        }
    }
}
