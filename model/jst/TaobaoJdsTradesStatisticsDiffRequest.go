package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单全链路状态统计差异比较 API请求
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

// 初始化TaobaoJdsTradesStatisticsDiffRequest对象
func NewTaobaoJdsTradesStatisticsDiffRequest() *TaobaoJdsTradesStatisticsDiffRequest{
    return &TaobaoJdsTradesStatisticsDiffRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJdsTradesStatisticsDiffRequest) GetApiMethodName() string {
    return "taobao.jds.trades.statistics.diff"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJdsTradesStatisticsDiffRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Date Setter
// 查询的日期，格式如YYYYMMDD的日期对应的数字
func (r *TaobaoJdsTradesStatisticsDiffRequest) SetDate(date int64) error {
    r.date = date
    r.Set("date", date)
    return nil
}

// Date Getter
func (r TaobaoJdsTradesStatisticsDiffRequest) GetDate() int64 {
    return r.date
}
// PreStatus Setter
// 需要比较的状态，将会和post_status做比较
func (r *TaobaoJdsTradesStatisticsDiffRequest) SetPreStatus(preStatus string) error {
    r.preStatus = preStatus
    r.Set("pre_status", preStatus)
    return nil
}

// PreStatus Getter
func (r TaobaoJdsTradesStatisticsDiffRequest) GetPreStatus() string {
    return r.preStatus
}
// PostStatus Setter
// 需要比较的状态
func (r *TaobaoJdsTradesStatisticsDiffRequest) SetPostStatus(postStatus string) error {
    r.postStatus = postStatus
    r.Set("post_status", postStatus)
    return nil
}

// PostStatus Getter
func (r TaobaoJdsTradesStatisticsDiffRequest) GetPostStatus() string {
    return r.postStatus
}
// PageNo Setter
// 页数
func (r *TaobaoJdsTradesStatisticsDiffRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoJdsTradesStatisticsDiffRequest) GetPageNo() int64 {
    return r.pageNo
}
