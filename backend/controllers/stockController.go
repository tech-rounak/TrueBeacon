package controllers

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	helper "github.com/tech-rounak/true-beacon/helpers"
	"github.com/tech-rounak/true-beacon/models"
)

func ParseCsvAndInsertIntoDB() error {

	//  if the data is already processed we don't proceed further
	var count int64
	db.Model(&models.Share{}).Count(&count)
	if count > 0 {
		return nil
	}

	fmt.Println("------------ Share Data Procsesssing --------------")
	file, err := os.Open("./historical_prices.csv")

	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		return err
	}
	for _, eachrecord := range records[1:] {
		layout := "2006-01-02"
		input := eachrecord[1][0:10]
		date, _ := time.Parse(layout, input)
		price, _ := strconv.ParseFloat(eachrecord[2], 64)
		fmt.Println(date, price, eachrecord[3])
		db.Create(&models.Share{
			Date:           date,
			Price:          price,
			InstrumentName: eachrecord[3],
		})
	}
	return nil
}

func FetchStockHistory() gin.HandlerFunc {
	return func(gc *gin.Context) {

		r := gc.Request

		symbol := r.URL.Query().Get("symbol")
		startDate := r.URL.Query().Get("startDate")
		endDate := r.URL.Query().Get("endDate")

		symbol = strings.ReplaceAll(symbol, "-", " ")
		layout := "2006-01-02"
		start, _ := time.Parse(layout, startDate)
		end, _ := time.Parse(layout, endDate)
		fmt.Println(start, end)
		rows, err := db.Raw("select * from shares where instrument_name=? AND date BETWEEN ? and ?", symbol, start, end).Rows()
		defer rows.Close()

		var share []models.Share
		if err != nil {
			fmt.Println("--------------- err ---------------", err)
			gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for rows.Next() {
			var tmpShare models.Share
			err := db.ScanRows(rows, &tmpShare)
			if err != nil {
				gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			share = append(share, tmpShare)

		}
		helper.Success(gc, 200, share, "Stock price history fetched successfully!!")
	}
}

// HandlerFunc to return the holdings in a portfolio
func PortfolioHoldings() gin.HandlerFunc {
	return func(gc *gin.Context) {
		res := map[string]interface{}{
			"holdings": []interface{}{
				map[string]interface{}{
					"tradingsymbol":         "GOLDBEES",
					"exchange":              "BSE",
					"isin":                  "INF204KB17I5",
					"quantity":              2,
					"authorised_date":       "2021-06-08 00:00:00",
					"average_price":         40.67,
					"last_price":            42.47,
					"close_price":           42.28,
					"pnl":                   3.60,
					"day_change":            0.19,
					"day_change_percentage": 0.45,
				},
				map[string]interface{}{
					"tradingsymbol":         "IDEA",
					"exchange":              "NSE",
					"isin":                  "INE669E01016",
					"quantity":              5,
					"authorised_date":       "2021-06-08 00:00:00",
					"average_price":         8.466,
					"last_price":            10.0,
					"close_price":           10.1,
					"pnl":                   7.67,
					"day_change":            -0.1,
					"day_change_percentage": -0.99,
				},
			},
		}
		helper.Success(gc, 200, res, "Holdings data fetched successfully!!")
	}
}

func PlaceOrder() gin.HandlerFunc {
	return func(gc *gin.Context) {
		rand.Seed(time.Now().UnixNano())

		randomInt := rand.Intn(1e9)

		data := map[string]interface{}{
			"order_id": randomInt,
		}
		helper.Success(gc, 200, data, "Order placed successfully!!")
	}
}
