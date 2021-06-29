package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询尖货活动信息 API请求
taobao.opentrade.activity.query

尖货交易活动信息配置，查询尖货活动信息
*/
type TaobaoOpentradeActivityQueryRequest struct {
    model.Params
    // 活动结束时间
    _endTime   string
    // 活动名称
    _activityName   string
    // 分页大小
    _pageSize   int64
    // 分页序号
    _pageIndex   int64
}

// 初始化TaobaoOpentradeActivityQueryRequest对象
func NewTaobaoOpentradeActivityQueryRequest() *TaobaoOpentradeActivityQueryRequest{
    return &TaobaoOpentradeActivityQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeActivityQueryRequest) GetApiMethodName() string {
    return "taobao.opentrade.activity.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeActivityQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EndTime Setter
// 活动结束时间
func (r *TaobaoOpentradeActivityQueryRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoOpentradeActivityQueryRequest) GetEndTime() string {
    return r._endTime
}
// ActivityName Setter
// 活动名称
func (r *TaobaoOpentradeActivityQueryRequest) SetActivityName(_activityName string) error {
    r._activityName = _activityName
    r.Set("activity_name", _activityName)
    return nil
}

// ActivityName Getter
func (r TaobaoOpentradeActivityQueryRequest) GetActivityName() string {
    return r._activityName
}
// PageSize Setter
// 分页大小
func (r *TaobaoOpentradeActivityQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpentradeActivityQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageIndex Setter
// 分页序号
func (r *TaobaoOpentradeActivityQueryRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoOpentradeActivityQueryRequest) GetPageIndex() int64 {
    return r._pageIndex
}
