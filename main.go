package main

// go:generate sqlboiler postgres

import (
	"database/sql"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/vattle/sqlboiler/boil"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

// Open handle to database like normal
var log = log15.New()

func main() {

	viper.SetConfigFile("./config.json")

	// Searches for config file in given paths and read it
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.SetDefault("POW", "http://api.f2pool.com/decred/address")
	viper.SetDefault("ExchangeData", "https://poloniex.com/public")

	db, err := sql.Open("postgres", "dbname="+viper.GetString("Database.pgdbname")+" user="+viper.GetString("Database.pguser")+"host="+viper.GetString("Database.pghost")+" password="+viper.GetString("Database.pgpass"))
	if err != nil {
		panic(err.Error())
		return
	}

	boil.SetDB(db)

	// fetchHistoricData(1, "") //parameters : exchangeID,currency pair, start time, end time
	getPOSdata()
	for {
		getHistoricData(1, "BTC-DCR", "1514764800", "1514851200") //parameters : exchangeID,currency pair, start time, end time                                    //parameters :  Currency pair
		getChartData(1, "BTC_DCR", "1405699200", "9999999999")    //parameters: exchange id,Currency Pair, start time , end time

	}

	getPOWData(2, "") //parameters: pool id

}

// Exchange id = 0 for Poloneix
// Exchange id =1 for Bittrex

// func fetchHistoricData(exchangeID int, date string) {

// 	if exchangeID == 0 { //Poloniex exchange
// 		user := Poloniex{

// 			client: &http.Client{},
// 		}
// 		user.fetchPoloniexData(date)

// 	}

// 	if exchangeID == 1 { //Bittrex Exchange

// 		user := Bittrex{

// 			client: &http.Client{},
// 		}
// 		user.fetchBittrexData(date)
// 	}

// }

func getPOSdata() {

	user := POS{
		client: &http.Client{},
	}

	user.getPOS()
}

func getPOWData(PoolID int, apiKey string) {

	user := POW{
		client: &http.Client{},
	}

	user.getPOW(PoolID, viper.GetString("POW"+"["+string(PoolID)+"]"), apiKey)

}

// Exchange id = 0 for Poloneix
// Exchange id =1 for Bittrex

func getHistoricData(exchangeID int, currencyPair string, startTime string, endTime string) {

	if exchangeID == 0 { //Poloniex exchange
		user := Poloniex{

			client: &http.Client{},
		}
		user.getPoloniexData(currencyPair, startTime, endTime)

	}

	if exchangeID == 1 { //Bittrex Exchange

		user := Bittrex{
			client: &http.Client{},
		}
		user.getBittrexData(currencyPair)
	}

	//Time delay of 24 hours

	time.Sleep(86400 * time.Second)
}

//Get chart data from exchanges
// if exchange id =0 , get chart data from poloniex exchange
// exchange id =1  get chart data from bittrex exchange

func getChartData(exchangeID int, currencyPair string, startTime string, endTime string) {

	if exchangeID == 0 {
		user := Poloniex{

			client: &http.Client{},
		}
		user.getChartData(currencyPair, startTime, endTime)

	}
	if exchangeID == 1 {
		user := Bittrex{
			client: &http.Client{},
		}
		user.getTicks(currencyPair)

	}

}
