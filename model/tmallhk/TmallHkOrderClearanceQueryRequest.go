package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫国际订单清关信息 APIRequest
tmall.hk.order.clearance.query

天猫国际订单清关信息查询
*/
type TmallHkOrderClearanceQueryRequest struct {
    model.Params

    // 交易主订单号
    bizOrderId   int64 

    // 调用方业务身份(由国际侧配置提供给调用方)
    businessSymbol   string 

}

func NewTmallHkOrderClearanceQueryRequest() *TmallHkOrderClearanceQueryRequest{
    return &TmallHkOrderClearanceQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallHkOrderClearanceQueryRequest) GetApiMethodName() string {
    return "tmall.hk.order.clearance.query"
}

func (r TmallHkOrderClearanceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallHkOrderClearanceQueryRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r TmallHkOrderClearanceQueryRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

func (r *TmallHkOrderClearanceQueryRequest) SetBusinessSymbol(businessSymbol string) error {
    r.businessSymbol = businessSymbol
    r.Set("business_symbol", businessSymbol)
    return nil
}

func (r TmallHkOrderClearanceQueryRequest) GetBusinessSymbol() string {
    return r.businessSymbol
}

