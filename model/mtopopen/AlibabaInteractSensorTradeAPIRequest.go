package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorTradeAPIRequest 交易组件 API请求
// alibaba.interact.sensor.trade
//
// 交易流程
type AlibabaInteractSensorTradeAPIRequest struct {
	model.Params
	// 系统自动生成
	_id string
}

// NewAlibabaInteractSensorTradeRequest 初始化AlibabaInteractSensorTradeAPIRequest对象
func NewAlibabaInteractSensorTradeRequest() *AlibabaInteractSensorTradeAPIRequest {
	return &AlibabaInteractSensorTradeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorTradeAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.trade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorTradeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetId is Id Setter
// 系统自动生成
func (r *AlibabaInteractSensorTradeAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaInteractSensorTradeAPIRequest) GetId() string {
	return r._id
}
