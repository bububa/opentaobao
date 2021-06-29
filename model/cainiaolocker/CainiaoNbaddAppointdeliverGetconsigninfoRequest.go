package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取支持定时派送服务发货信息 API请求
cainiao.nbadd.appointdeliver.getconsigninfo

获取支持定时派送服务发货信息
*/
type CainiaoNbaddAppointdeliverGetconsigninfoRequest struct {
    model.Params
    // 淘宝交易订单id
    _tradeOrderId   int64
}

// 初始化CainiaoNbaddAppointdeliverGetconsigninfoRequest对象
func NewCainiaoNbaddAppointdeliverGetconsigninfoRequest() *CainiaoNbaddAppointdeliverGetconsigninfoRequest{
    return &CainiaoNbaddAppointdeliverGetconsigninfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoNbaddAppointdeliverGetconsigninfoRequest) GetApiMethodName() string {
    return "cainiao.nbadd.appointdeliver.getconsigninfo"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoNbaddAppointdeliverGetconsigninfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeOrderId Setter
// 淘宝交易订单id
func (r *CainiaoNbaddAppointdeliverGetconsigninfoRequest) SetTradeOrderId(_tradeOrderId int64) error {
    r._tradeOrderId = _tradeOrderId
    r.Set("trade_order_id", _tradeOrderId)
    return nil
}

// TradeOrderId Getter
func (r CainiaoNbaddAppointdeliverGetconsigninfoRequest) GetTradeOrderId() int64 {
    return r._tradeOrderId
}
