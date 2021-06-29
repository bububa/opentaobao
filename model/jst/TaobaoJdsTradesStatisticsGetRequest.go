package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单数量统计结果 API请求
taobao.jds.trades.statistics.get

获取订单数量统计结果
*/
type TaobaoJdsTradesStatisticsGetRequest struct {
    model.Params
    // 查询的日期，格式如YYYYMMDD的日期对应的数字
    _date   int64
}

// 初始化TaobaoJdsTradesStatisticsGetRequest对象
func NewTaobaoJdsTradesStatisticsGetRequest() *TaobaoJdsTradesStatisticsGetRequest{
    return &TaobaoJdsTradesStatisticsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJdsTradesStatisticsGetRequest) GetApiMethodName() string {
    return "taobao.jds.trades.statistics.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJdsTradesStatisticsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Date Setter
// 查询的日期，格式如YYYYMMDD的日期对应的数字
func (r *TaobaoJdsTradesStatisticsGetRequest) SetDate(_date int64) error {
    r._date = _date
    r.Set("date", _date)
    return nil
}

// Date Getter
func (r TaobaoJdsTradesStatisticsGetRequest) GetDate() int64 {
    return r._date
}
