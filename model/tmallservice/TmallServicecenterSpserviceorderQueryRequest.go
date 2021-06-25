package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务单列表查询 APIRequest
tmall.servicecenter.spserviceorder.query

查询服务单列表
*/
type TmallServicecenterSpserviceorderQueryRequest struct {
    model.Params

    // 交易主订单id
    parentBizOrderId   int64 

}

func NewTmallServicecenterSpserviceorderQueryRequest() *TmallServicecenterSpserviceorderQueryRequest{
    return &TmallServicecenterSpserviceorderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterSpserviceorderQueryRequest) GetApiMethodName() string {
    return "tmall.servicecenter.spserviceorder.query"
}

func (r TmallServicecenterSpserviceorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterSpserviceorderQueryRequest) SetParentBizOrderId(parentBizOrderId int64) error {
    r.parentBizOrderId = parentBizOrderId
    r.Set("parent_biz_order_id", parentBizOrderId)
    return nil
}

func (r TmallServicecenterSpserviceorderQueryRequest) GetParentBizOrderId() int64 {
    return r.parentBizOrderId
}

