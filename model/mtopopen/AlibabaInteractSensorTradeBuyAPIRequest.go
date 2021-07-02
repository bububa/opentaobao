package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorTradeBuyAPIRequest 手淘下单能力开放 API请求
// alibaba.interact.sensor.trade.buy
//
// 交易流程鉴权
type AlibabaInteractSensorTradeBuyAPIRequest struct {
	model.Params
	// 系统自动生成
	_id string
}

// NewAlibabaInteractSensorTradeBuyRequest 初始化AlibabaInteractSensorTradeBuyAPIRequest对象
func NewAlibabaInteractSensorTradeBuyRequest() *AlibabaInteractSensorTradeBuyAPIRequest {
	return &AlibabaInteractSensorTradeBuyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorTradeBuyAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.trade.buy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorTradeBuyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Id Setter
// 系统自动生成
func (r *AlibabaInteractSensorTradeBuyAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r AlibabaInteractSensorTradeBuyAPIRequest) GetId() string {
	return r._id
}
