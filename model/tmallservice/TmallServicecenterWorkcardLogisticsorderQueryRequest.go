package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流单信息查询 APIRequest
tmall.servicecenter.workcard.logisticsorder.query

物流订单信息查询API
*/
type TmallServicecenterWorkcardLogisticsorderQueryRequest struct {
    model.Params

    // 物流单号
    logisticsOrderId   int64 

}

func NewTmallServicecenterWorkcardLogisticsorderQueryRequest() *TmallServicecenterWorkcardLogisticsorderQueryRequest{
    return &TmallServicecenterWorkcardLogisticsorderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardLogisticsorderQueryRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.logisticsorder.query"
}

func (r TmallServicecenterWorkcardLogisticsorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardLogisticsorderQueryRequest) SetLogisticsOrderId(logisticsOrderId int64) error {
    r.logisticsOrderId = logisticsOrderId
    r.Set("logistics_order_id", logisticsOrderId)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderQueryRequest) GetLogisticsOrderId() int64 {
    return r.logisticsOrderId
}

