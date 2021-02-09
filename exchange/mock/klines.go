package mock

import (
	"github.com/long2ice/trader/db"
	"github.com/long2ice/trader/exchange"
)

type KLineService struct {
	exchange.KLineService
	Api       exchange.IApi
	Symbol    string
	Interval  string
	StartTime int
	EndTime   int
	Limit     int
}

func (service *KLineService) Do() ([]exchange.KLine, error) {
	var kLines []db.KLine
	db.Client.Where("symbol = ?", service.Symbol).Where("open_time >= ?", service.StartTime).Limit(service.Limit).Order("open_time").Find(&kLines)
	var ret []exchange.KLine
	for _, line := range kLines {
		ret = append(ret, exchange.KLine{
			Open:      line.Open,
			Close:     line.Close,
			High:      line.High,
			Low:       line.Low,
			Amount:    line.Amount,
			Volume:    line.Vol,
			Finish:    true,
			CloseTime: line.CloseTime,
		})
	}
	return ret, nil
}
