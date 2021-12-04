# inner

Get the inner structure of any value(interface{}).


## Getting Started

### Install

`go get https://github.com/i0Ek3/inner`

### Usage

```Go
package main

import (
    "github.com/i0Ek3/inner"
)

func main() {
    yogurt := Item{
        Name:    "Kight Yogurt",
        ID:      "1520012728X",
        Year:    2021,
        Month:   11,
        Date:    29,
        Stocked: true,
        Where: map[string]string{
            "Product X": "IM",
            "Product Y": "LN",
        },
        Compentent: []string{
            "Raw milk",
            "White sugar",
            "Minneral water",
            "Edible flavor",
        },
    }

    // Inner gets the structure inner of x
    inner.Inner("yogurt", yogurt)
    /* output:
    ["yogurt" => inner.Item]
    yogurt.Name = "Kight Yogurt"
    yogurt.ID = "1520012728X"
    yogurt.Year = 2021
    yogurt.Month = 11
    yogurt.Date = 29
    yogurt.Stocked = true
    yogurt.Where["Product X"] = "IM"
    yogurt.Where["Product Y"] = "LN"
    yogurt.Compentent[0] = "Raw milk"
    yogurt.Compentent[1] = "White sugar"
    yogurt.Compentent[2] = "Minneral water"
    yogurt.Compentent[3] = "Edible flavor"
    */

}
```

## TODO

- improve test file
- expand type process
- bug fix


## Credit

- gopl.io
- @wxnacy
