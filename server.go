package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/calculate_tax", calculateTax)
	e.Logger.Fatal(e.Start(":1323"))
}

// Response - Tax response from server
type Response struct {
	WithOutTax float64 `json:"withOutTax"`
	TaxRate    float64 `json:"taxRate"`
	Amount     float64 `json:"amount"`
	TaxAmount  float64 `json:"TaxAmount"`
}

func calculateTax(c echo.Context) error {
	tax := c.QueryParam("tax")
	amount := c.QueryParam("amount")

	r := &Response{
		WithOutTax: withOutTax(amount, tax),
		TaxRate:    convertTax(tax),
		Amount:     convertAmoutn(amount),
		TaxAmount:  calculateTaxAmount(amount, tax),
	}
	if err := c.Bind(r); err != nil {
		return err
	}
	return c.JSONPretty(http.StatusCreated, r, " ")
}

func withOutTax(amount, taxRate string) float64 {
	intAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println(err)
	}
	intTaxRate, err := strconv.ParseFloat(taxRate, 64)
	if err != nil {
		fmt.Println(err)
	}
	taxAmount := intAmount / (1.0 + intTaxRate)
	return taxAmount
}

func calculateTaxAmount(amount, taxRate string) float64 {
	intAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println(err)
	}
	intTaxRate, err := strconv.ParseFloat(taxRate, 64)
	if err != nil {
		fmt.Println(err)
	}
	taxAmount := intAmount / (1.0 + intTaxRate)
	return intAmount - taxAmount
}

func convertTax(tr string) float64 {
	itr, err := strconv.ParseFloat(tr, 64)
	if err != nil {
		log.Fatal(err)
	}
	return itr
}

func convertAmoutn(a string) float64 {
	ia, err := strconv.ParseFloat(a, 64)
	if err != nil {
		log.Fatal(err)
	}
	return ia
}
