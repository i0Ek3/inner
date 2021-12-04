package main

import (
    "fmt"

    "github.com/i0Ek3/inner"
)

type Item struct {
    Name, ID          string
    Year, Month, Date int
    Stocked           bool
    Where             map[string]string
    Compentent        []string
}

var (
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

    sausage := Item{
        Name:    "Italian Sausage",
        ID:      "1520012728Y",
        Year:    2021,
        Month:   12,
        Date:    "03", // wrong fromat
        Stocked: true,
        Where: map[string]string{
            "Product X": "CN",
            "Product Y": "IT",
        },
        Compentent: []string{
            "Pork",
            "White sugar",
            "Minneral water",
            "Edible flavor",
        },
    }
)

func main() {
    inner.Inner("yogurt", yogurt)
    inner.Inner("sausage", sausage)
}
