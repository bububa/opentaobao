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
    endTime   string
    // 活动名称
    activityName   string
    // 分页大小
    pageSize   int64
    // 分页序号
    pageIndex   int64
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
func (r *TaobaoOpentradeActivityQueryRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoOpentradeActivityQueryRequest) GetEndTime() string {
    return r.endTime
}
// ActivityName Setter
// 活动名称
func (r *TaobaoOpentradeActivityQueryRequest) SetActivityName(activityName string) error {
    r.activityName = activityName
    r.Set("activity_name", activityName)
    return nil
}

// ActivityName Getter
func (r TaobaoOpentradeActivityQueryRequest) GetActivityName() string {
    return r.activityName
}
// PageSize Setter
// 分页大小
func (r *TaobaoOpentradeActivityQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpentradeActivityQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageIndex Setter
// 分页序号
func (r *TaobaoOpentradeActivityQueryRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoOpentradeActivityQueryRequest) GetPageIndex() int64 {
    return r.pageIndex
}
