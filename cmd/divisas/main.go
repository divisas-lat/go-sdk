package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/divisas-lat/go-sdk"
	"github.com/divisas-lat/go-sdk/enums"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	client := divisas.NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	switch command {
	case "today":
		handleToday(ctx, client)
	case "convert":
		handleConvert(ctx, client)
	case "stats":
		handleStats(ctx, client)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: divisas <command> [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  today <country_code> [currency_code]")
	fmt.Println("  convert <amount> <from_currency> to <to_currency> in <country_code>")
	fmt.Println("  stats <country_code> [period]")
}

func handleToday(ctx context.Context, client *divisas.Client) {
	cmd := flag.NewFlagSet("today", flag.ExitOnError)
	cmd.Parse(os.Args[2:])
	args := cmd.Args()

	if len(args) < 1 {
		fmt.Println("Usage: divisas today <country_code> [currency_code]")
		os.Exit(1)
	}

	country := enums.Country(strings.ToUpper(args[0]))
	q := client.Query().ForCountry(country)

	if len(args) >= 2 {
		q = q.WithCurrency(enums.Currency(strings.ToUpper(args[1])))
	}

	res, err := q.GetToday(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Country: %s | Base: %s | Date: %s\n", res.Country, res.BaseCurrency, res.Date)
	if res.Rate != nil {
		fmt.Printf("Rate: %s - Buy: %.4f / Sell: %.4f\n", res.Rate.CurrencyCode, res.Rate.Buy, res.Rate.Sell)
	} else if len(res.Rates) > 0 {
		for _, r := range res.Rates {
			fmt.Printf("Rate: %s - Buy: %.4f / Sell: %.4f\n", r.CurrencyCode, r.Buy, r.Sell)
		}
	}
}

func handleConvert(ctx context.Context, client *divisas.Client) {
	cmd := flag.NewFlagSet("convert", flag.ExitOnError)
	cmd.Parse(os.Args[2:])
	args := cmd.Args()

	// convert 100 USD to GTQ in GT
	// indices: 0: 100, 1: USD, 2: to, 3: GTQ, 4: in, 5: GT
	if len(args) < 6 || args[2] != "to" || args[4] != "in" {
		fmt.Println("Usage: divisas convert <amount> <from_currency> to <to_currency> in <country_code>")
		os.Exit(1)
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Printf("Invalid amount: %v\n", err)
		os.Exit(1)
	}

	fromCurrency := enums.Currency(strings.ToUpper(args[1]))
	toCurrency := enums.Currency(strings.ToUpper(args[3]))
	country := enums.Country(strings.ToUpper(args[5]))

	res, err := client.Query().
		ForCountry(country).
		WithCurrency(fromCurrency).
		Convert(ctx, toCurrency, amount)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %.2f %s = %.2f %s\n", res.From.Amount, res.From.Currency, res.Result, res.To.Currency)
	fmt.Printf("Effective Rate: %.4f | Via: %s\n", res.EffectiveRate, res.Via)
}

func handleStats(ctx context.Context, client *divisas.Client) {
	cmd := flag.NewFlagSet("stats", flag.ExitOnError)
	cmd.Parse(os.Args[2:])
	args := cmd.Args()

	if len(args) < 1 {
		fmt.Println("Usage: divisas stats <country_code> [period]")
		os.Exit(1)
	}

	country := enums.Country(strings.ToUpper(args[0]))
	period := "30d"
	if len(args) >= 2 {
		period = args[1]
	}

	res, err := client.Query().
		ForCountry(country).
		GetStats(ctx, period)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Stats for %s (%s)\n", res.Country, res.Period)
	fmt.Printf("Min Buy: %.4f | Max Buy: %.4f | Avg Buy: %.4f\n", res.Stats.MinBuy, res.Stats.MaxBuy, res.Stats.AvgBuy)
	fmt.Printf("Volatility: %.4f%%\n", res.Stats.Volatility)
}
