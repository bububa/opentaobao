package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按照日期范围查询物流订单详情 APIRequest
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

func NewTaobaoWlbOrderdetailDateGetRequest() *TaobaoWlbOrderdetailDateGetRequest{
    return &TaobaoWlbOrderdetailDateGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbOrderdetailDateGetRequest) GetApiMethodName() string {
    return "taobao.wlb.orderdetail.date.get"
}

func (r TaobaoWlbOrderdetailDateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbOrderdetailDateGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoWlbOrderdetailDateGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoWlbOrderdetailDateGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoWlbOrderdetailDateGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoWlbOrderdetailDateGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoWlbOrderdetailDateGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoWlbOrderdetailDateGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoWlbOrderdetailDateGetRequest) GetPageNo() int64 {
    return r.pageNo
}

