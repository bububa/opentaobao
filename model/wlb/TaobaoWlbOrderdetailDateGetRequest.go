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
    startTime   string
    // 创建时间结束
    endTime   string
    // 分页大小
    pageSize   int64
    // 分页下标
    pageNo   int64
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
func (r *TaobaoWlbOrderdetailDateGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoWlbOrderdetailDateGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 创建时间结束
func (r *TaobaoWlbOrderdetailDateGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoWlbOrderdetailDateGetRequest) GetEndTime() string {
    return r.endTime
}
// PageSize Setter
// 分页大小
func (r *TaobaoWlbOrderdetailDateGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoWlbOrderdetailDateGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNo Setter
// 分页下标
func (r *TaobaoWlbOrderdetailDateGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoWlbOrderdetailDateGetRequest) GetPageNo() int64 {
    return r.pageNo
}
