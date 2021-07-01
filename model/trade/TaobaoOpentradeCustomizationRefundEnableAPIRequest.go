package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定制订单设置允许仅退款 API请求
taobao.opentrade.customization.refund.enable

定制订单设置允许仅退款
*/
type TaobaoOpentradeCustomizationRefundEnableAPIRequest struct {
    model.Params
    // 主订单ID
    _tradeId   int64
}

// 初始化TaobaoOpentradeCustomizationRefundEnableAPIRequest对象
func NewTaobaoOpentradeCustomizationRefundEnableRequest() *TaobaoOpentradeCustomizationRefundEnableAPIRequest{
    return &TaobaoOpentradeCustomizationRefundEnableAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeCustomizationRefundEnableAPIRequest) GetApiMethodName() string {
    return "taobao.opentrade.customization.refund.enable"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeCustomizationRefundEnableAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeId Setter
// 主订单ID
func (r *TaobaoOpentradeCustomizationRefundEnableAPIRequest) SetTradeId(_tradeId int64) error {
    r._tradeId = _tradeId
    r.Set("trade_id", _tradeId)
    return nil
}

// TradeId Getter
func (r TaobaoOpentradeCustomizationRefundEnableAPIRequest) GetTradeId() int64 {
    return r._tradeId
}
