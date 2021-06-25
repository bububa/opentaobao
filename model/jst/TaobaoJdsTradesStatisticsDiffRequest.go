package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单全链路状态统计差异比较 APIRequest
taobao.jds.trades.statistics.diff

订单全链路状态统计差异比较
*/
type TaobaoJdsTradesStatisticsDiffRequest struct {
    model.Params

    // 查询的日期，格式如YYYYMMDD的日期对应的数字
    date   int64 

    // 需要比较的状态，将会和post_status做比较
    preStatus   string 

    // 需要比较的状态
    postStatus   string 

    // 页数
    pageNo   int64 

}

func NewTaobaoJdsTradesStatisticsDiffRequest() *TaobaoJdsTradesStatisticsDiffRequest{
    return &TaobaoJdsTradesStatisticsDiffRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJdsTradesStatisticsDiffRequest) GetApiMethodName() string {
    return "taobao.jds.trades.statistics.diff"
}

func (r TaobaoJdsTradesStatisticsDiffRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJdsTradesStatisticsDiffRequest) SetDate(date int64) error {
    r.date = date
    r.Set("date", date)
    return nil
}

func (r TaobaoJdsTradesStatisticsDiffRequest) GetDate() int64 {
    return r.date
}

func (r *TaobaoJdsTradesStatisticsDiffRequest) SetPreStatus(preStatus string) error {
    r.preStatus = preStatus
    r.Set("pre_status", preStatus)
    return nil
}

func (r TaobaoJdsTradesStatisticsDiffRequest) GetPreStatus() string {
    return r.preStatus
}

func (r *TaobaoJdsTradesStatisticsDiffRequest) SetPostStatus(postStatus string) error {
    r.postStatus = postStatus
    r.Set("post_status", postStatus)
    return nil
}

func (r TaobaoJdsTradesStatisticsDiffRequest) GetPostStatus() string {
    return r.postStatus
}

func (r *TaobaoJdsTradesStatisticsDiffRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoJdsTradesStatisticsDiffRequest) GetPageNo() int64 {
    return r.pageNo
}

