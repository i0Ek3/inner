package inner

import (
    "testing"
)

type Item struct {
    Name, ID          string
    Year, Month, Date int
    Stocked           bool
    Where             map[string]string
    Compentent        []string
}

var (
    yogurt = Item{
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

    sausage = Item{
        Name:    "Italian Sausage",
        ID:      "1520012728Y",
        Year:    2021,
        Month:   12,
        Date:    03,
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

func TestInner(t *testing.T) {
    Inner("yogurt", yogurt)
    Inner("sausage", sausage)
}

/*
func ExampleInner() {
    Inner("yogurt", yogurt)
    // output:
    // yogurt.Name = "Kight Yogurt"
    // yogurt.ID = "1520012728X"
    // yogurt.Year = 2021
    // yogurt.Month = 11
    // yogurt.Date = 29
    // yogurt.Stocked = true
    // yogurt.Where["Product X"] = "IM"
    // yogurt.Where["Product Y"] = "LN"
    // yogurt.Compentent[0] = "Raw milk"
    // yogurt.Compentent[1] = "White sugar"
    // yogurt.Compentent[2] = "Minneral water"
    // yogurt.Compentent[3] = "Edible flavor"
}*/

func TestAny(t *testing.T) {
    var i Item
    Anyln(yogurt)
    Anyln(sausage)
    Anyln(i.ID)
    Anyln(i.Year)
    Anyln(i.Where)
    Anyln(i.Compentent)
    Anyln(i.Stocked)
    var ch chan int
    Anyln(ch)
}
