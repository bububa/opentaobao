package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorgravityAPIRequest 重力感应 API请求
// alibaba.interact.sensor.gravity
//
// native获取重力感应
type AlibabainteractsensorgravityAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorgravityRequest 初始化AlibabainteractsensorgravityAPIRequest对象
func NewAlibabainteractsensorgravityRequest() *AlibabainteractsensorgravityAPIRequest {
	return &AlibabainteractsensorgravityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorgravityAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.gravity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorgravityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorgravityAPIRequest) GetRawParams() model.Params {
	return r.Params
}
