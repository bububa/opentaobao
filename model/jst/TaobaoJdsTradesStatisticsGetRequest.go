package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单数量统计结果 APIRequest
taobao.jds.trades.statistics.get

获取订单数量统计结果
*/
type TaobaoJdsTradesStatisticsGetRequest struct {
    model.Params

    // 查询的日期，格式如YYYYMMDD的日期对应的数字
    date   int64 

}

func NewTaobaoJdsTradesStatisticsGetRequest() *TaobaoJdsTradesStatisticsGetRequest{
    return &TaobaoJdsTradesStatisticsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJdsTradesStatisticsGetRequest) GetApiMethodName() string {
    return "taobao.jds.trades.statistics.get"
}

func (r TaobaoJdsTradesStatisticsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJdsTradesStatisticsGetRequest) SetDate(date int64) error {
    r.date = date
    r.Set("date", date)
    return nil
}

func (r TaobaoJdsTradesStatisticsGetRequest) GetDate() int64 {
    return r.date
}

