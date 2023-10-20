package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensortradeAPIRequest 交易组件 API请求
// alibaba.interact.sensor.trade
//
// 交易流程
type AlibabainteractsensortradeAPIRequest struct {
	model.Params
	// 系统自动生成
	_id string
}

// NewAlibabainteractsensortradeRequest 初始化AlibabainteractsensortradeAPIRequest对象
func NewAlibabainteractsensortradeRequest() *AlibabainteractsensortradeAPIRequest {
	return &AlibabainteractsensortradeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensortradeAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.trade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensortradeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensortradeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 系统自动生成
func (r *AlibabainteractsensortradeAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabainteractsensortradeAPIRequest) GetId() string {
	return r._id
}
