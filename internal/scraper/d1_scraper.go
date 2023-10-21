package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/MontFerret/ferret/pkg/compiler"
	"github.com/MontFerret/ferret/pkg/drivers"
	"github.com/MontFerret/ferret/pkg/drivers/cdp"
)

type D1Product struct {
    Name string `json:"product_name"`
}

func Scrape() {
    body, err := os.ReadFile("./internal/queries/d1/by_category.fql")
    if err != nil {
        panic(err)
    }

    query := string(body)

    comp := compiler.New()
    program, err := comp.Compile(query)

    if err != nil {
        panic(err)
    }

    ctx := context.Background()

	ctx = drivers.WithContext(ctx, cdp.NewDriver(), drivers.AsDefault())

	out, err := program.Run(ctx)

	if err != nil {
		panic(err)
	}

	res := make([]*D1Product, 0, 10)

    fmt.Println(string(out))

	err = json.Unmarshal(out, &res)

	if err != nil {
		panic(err)
	}

    fmt.Println(res)
	for _, topic := range res {
		fmt.Println(topic)
	}
}
