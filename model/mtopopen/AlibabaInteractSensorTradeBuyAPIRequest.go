package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensortradebuyAPIRequest 手淘下单能力开放 API请求
// alibaba.interact.sensor.trade.buy
//
// 交易流程鉴权
type AlibabainteractsensortradebuyAPIRequest struct {
	model.Params
	// 系统自动生成
	_id string
}

// NewAlibabainteractsensortradebuyRequest 初始化AlibabainteractsensortradebuyAPIRequest对象
func NewAlibabainteractsensortradebuyRequest() *AlibabainteractsensortradebuyAPIRequest {
	return &AlibabainteractsensortradebuyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensortradebuyAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.trade.buy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensortradebuyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensortradebuyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 系统自动生成
func (r *AlibabainteractsensortradebuyAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabainteractsensortradebuyAPIRequest) GetId() string {
	return r._id
}
