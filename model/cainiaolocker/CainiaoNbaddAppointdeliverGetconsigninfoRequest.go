package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取支持定时派送服务发货信息 APIRequest
cainiao.nbadd.appointdeliver.getconsigninfo

获取支持定时派送服务发货信息
*/
type CainiaoNbaddAppointdeliverGetconsigninfoRequest struct {
    model.Params

    // 淘宝交易订单id
    tradeOrderId   int64 

}

func NewCainiaoNbaddAppointdeliverGetconsigninfoRequest() *CainiaoNbaddAppointdeliverGetconsigninfoRequest{
    return &CainiaoNbaddAppointdeliverGetconsigninfoRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoNbaddAppointdeliverGetconsigninfoRequest) GetApiMethodName() string {
    return "cainiao.nbadd.appointdeliver.getconsigninfo"
}

func (r CainiaoNbaddAppointdeliverGetconsigninfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoNbaddAppointdeliverGetconsigninfoRequest) SetTradeOrderId(tradeOrderId int64) error {
    r.tradeOrderId = tradeOrderId
    r.Set("trade_order_id", tradeOrderId)
    return nil
}

func (r CainiaoNbaddAppointdeliverGetconsigninfoRequest) GetTradeOrderId() int64 {
    return r.tradeOrderId
}

