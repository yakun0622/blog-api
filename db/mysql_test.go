package db

import (
	"testing"
	"time"
)

type BtcUsdModel1 struct {
	id     uint64    `json:"id"`
	time   time.Time `json:"time"`
	open   float64   `json:"open"`
	high   float64   `json:"high"`
	low    float64   `json:"low"`
	close  float64   `json:"close"`
	volume float64   `json:"volume"`
}

func TestFetchRows(t *testing.T) {
	sql := "select * from btc_usd_15m limit 0, 10"
	rows, err := FetchRows(sql)
	if err != nil {
		t.Error(err)
		return
	}
	defer rows.Close()

	datas := make([]BtcUsdModel1, 0)
	for rows.Next() {
		var data BtcUsdModel1
		rows.Scan(&data.id, &data.time, &data.open, &data.high, &data.low, &data.close, &data.volume)
		datas = append(datas, data)
	}

	t.Log(datas)
}
