package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预售结算数据回传 APIRequest
alibaba.mj.presale.settlement.addlist

用于预售活动结算数据的回传。
*/
type AlibabaMjPresaleSettlementAddlistRequest struct {
    model.Params

    // 订单json格式数据
    preSaleRefundJson   string 

}

func NewAlibabaMjPresaleSettlementAddlistRequest() *AlibabaMjPresaleSettlementAddlistRequest{
    return &AlibabaMjPresaleSettlementAddlistRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjPresaleSettlementAddlistRequest) GetApiMethodName() string {
    return "alibaba.mj.presale.settlement.addlist"
}

func (r AlibabaMjPresaleSettlementAddlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjPresaleSettlementAddlistRequest) SetPreSaleRefundJson(preSaleRefundJson string) error {
    r.preSaleRefundJson = preSaleRefundJson
    r.Set("pre_sale_refund_json", preSaleRefundJson)
    return nil
}

func (r AlibabaMjPresaleSettlementAddlistRequest) GetPreSaleRefundJson() string {
    return r.preSaleRefundJson
}

