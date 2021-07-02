package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoNbaddAppointdeliverGetconsigninfoAPIRequest 获取支持定时派送服务发货信息 API请求
// cainiao.nbadd.appointdeliver.getconsigninfo
//
// 获取支持定时派送服务发货信息
type CainiaoNbaddAppointdeliverGetconsigninfoAPIRequest struct {
	model.Params
	// 淘宝交易订单id
	_tradeOrderId int64
}

// NewCainiaoNbaddAppointdeliverGetconsigninfoRequest 初始化CainiaoNbaddAppointdeliverGetconsigninfoAPIRequest对象
func NewCainiaoNbaddAppointdeliverGetconsigninfoRequest() *CainiaoNbaddAppointdeliverGetconsigninfoAPIRequest {
	return &CainiaoNbaddAppointdeliverGetconsigninfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoNbaddAppointdeliverGetconsigninfoAPIRequest) GetApiMethodName() string {
	return "cainiao.nbadd.appointdeliver.getconsigninfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoNbaddAppointdeliverGetconsigninfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTradeOrderId is TradeOrderId Setter
// 淘宝交易订单id
func (r *CainiaoNbaddAppointdeliverGetconsigninfoAPIRequest) SetTradeOrderId(_tradeOrderId int64) error {
	r._tradeOrderId = _tradeOrderId
	r.Set("trade_order_id", _tradeOrderId)
	return nil
}

// GetTradeOrderId TradeOrderId Getter
func (r CainiaoNbaddAppointdeliverGetconsigninfoAPIRequest) GetTradeOrderId() int64 {
	return r._tradeOrderId
}
