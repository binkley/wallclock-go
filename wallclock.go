package wallclock

import (
    "fmt"
    "strings"
)

func Minutes(minute int) ([]string, error) {
    switch(minute) {
    case 5:
        return strings.Fields("five past"), nil
    }
    return nil, fmt.Errorf("wallclock.Minutes: out of range: %d", minute)
}
