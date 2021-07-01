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
type TaobaoJdsTradesStatisticsDiffAPIRequest struct {
    model.Params
    // 查询的日期，格式如YYYYMMDD的日期对应的数字
    _date   int64
    // 需要比较的状态，将会和post_status做比较
    _preStatus   string
    // 需要比较的状态
    _postStatus   string
    // 页数
    _pageNo   int64
}

// 初始化TaobaoJdsTradesStatisticsDiffAPIRequest对象
func NewTaobaoJdsTradesStatisticsDiffRequest() *TaobaoJdsTradesStatisticsDiffAPIRequest{
    return &TaobaoJdsTradesStatisticsDiffAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJdsTradesStatisticsDiffAPIRequest) GetApiMethodName() string {
    return "taobao.jds.trades.statistics.diff"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJdsTradesStatisticsDiffAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Date Setter
// 查询的日期，格式如YYYYMMDD的日期对应的数字
func (r *TaobaoJdsTradesStatisticsDiffAPIRequest) SetDate(_date int64) error {
    r._date = _date
    r.Set("date", _date)
    return nil
}

// Date Getter
func (r TaobaoJdsTradesStatisticsDiffAPIRequest) GetDate() int64 {
    return r._date
}
// PreStatus Setter
// 需要比较的状态，将会和post_status做比较
func (r *TaobaoJdsTradesStatisticsDiffAPIRequest) SetPreStatus(_preStatus string) error {
    r._preStatus = _preStatus
    r.Set("pre_status", _preStatus)
    return nil
}

// PreStatus Getter
func (r TaobaoJdsTradesStatisticsDiffAPIRequest) GetPreStatus() string {
    return r._preStatus
}
// PostStatus Setter
// 需要比较的状态
func (r *TaobaoJdsTradesStatisticsDiffAPIRequest) SetPostStatus(_postStatus string) error {
    r._postStatus = _postStatus
    r.Set("post_status", _postStatus)
    return nil
}

// PostStatus Getter
func (r TaobaoJdsTradesStatisticsDiffAPIRequest) GetPostStatus() string {
    return r._postStatus
}
// PageNo Setter
// 页数
func (r *TaobaoJdsTradesStatisticsDiffAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoJdsTradesStatisticsDiffAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
