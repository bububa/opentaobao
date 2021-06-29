package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按照日期范围查询物流订单详情 API请求
taobao.wlb.orderdetail.date.get

外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情
*/
type TaobaoWlbOrderdetailDateGetRequest struct {
    model.Params
    // 创建时间起始
    _startTime   string
    // 创建时间结束
    _endTime   string
    // 分页大小
    _pageSize   int64
    // 分页下标
    _pageNo   int64
}

// 初始化TaobaoWlbOrderdetailDateGetRequest对象
func NewTaobaoWlbOrderdetailDateGetRequest() *TaobaoWlbOrderdetailDateGetRequest{
    return &TaobaoWlbOrderdetailDateGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderdetailDateGetRequest) GetApiMethodName() string {
    return "taobao.wlb.orderdetail.date.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderdetailDateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartTime Setter
// 创建时间起始
func (r *TaobaoWlbOrderdetailDateGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoWlbOrderdetailDateGetRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 创建时间结束
func (r *TaobaoWlbOrderdetailDateGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoWlbOrderdetailDateGetRequest) GetEndTime() string {
    return r._endTime
}
// PageSize Setter
// 分页大小
func (r *TaobaoWlbOrderdetailDateGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoWlbOrderdetailDateGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 分页下标
func (r *TaobaoWlbOrderdetailDateGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoWlbOrderdetailDateGetRequest) GetPageNo() int64 {
    return r._pageNo
}
