package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定制订单设置允许仅退款 APIRequest
taobao.opentrade.customization.refund.enable

定制订单设置允许仅退款
*/
type TaobaoOpentradeCustomizationRefundEnableRequest struct {
    model.Params

    // 主订单ID
    tradeId   int64 

}

func NewTaobaoOpentradeCustomizationRefundEnableRequest() *TaobaoOpentradeCustomizationRefundEnableRequest{
    return &TaobaoOpentradeCustomizationRefundEnableRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeCustomizationRefundEnableRequest) GetApiMethodName() string {
    return "taobao.opentrade.customization.refund.enable"
}

func (r TaobaoOpentradeCustomizationRefundEnableRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeCustomizationRefundEnableRequest) SetTradeId(tradeId int64) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

func (r TaobaoOpentradeCustomizationRefundEnableRequest) GetTradeId() int64 {
    return r.tradeId
}

