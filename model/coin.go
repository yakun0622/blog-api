package model

import (
	"blog_api/db"
	"blog_api/log"
	"strconv"
	"time"
)

type Coin struct {
	Id     uint64    `json:"id"`
	Time   time.Time `json:"time"`
	Open   float32   `json:"open"`
	High   float32   `json:"high"`
	Low    float32   `json:"low"`
	Close  float32   `json:"close"`
	Volume float32   `json:"volume"`
}

func (coin *Coin) QueryList(tableName string) ([]Coin, [][]string, error) {
	sql := "select id,time, open, high, low, close, volume from " + tableName + " limit 0, 1000"
	rows, err := db.FetchRows(sql)
	if err != nil {
		log.Logger.Error(err)
		return nil, nil, err
	}
	defer rows.Close()

	datas := make([]Coin, 0)
	var dataArray [][]string
	for rows.Next() {
		data := Coin{}
		rows.Scan(&data.Id, &data.Time, &data.Open, &data.High, &data.Low, &data.Close, &data.Volume)
		datas = append(datas, data)
		var item []string
		item = append(item, data.Time.Format("2006-01-02 15:04:05"))
		item = append(item, strconv.FormatFloat(float64(data.Open), 'f', 6, 64))
		item = append(item, strconv.FormatFloat(float64(data.High), 'f', 6, 64))
		item = append(item, strconv.FormatFloat(float64(data.Low), 'f', 6, 64))
		item = append(item, strconv.FormatFloat(float64(data.Close), 'f', 6, 64))
		dataArray = append(dataArray, item)
	}

	if err = rows.Err(); err != nil {
		log.Logger.Fatal(err)
	}

	return datas, dataArray, nil
}
