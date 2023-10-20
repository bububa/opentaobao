package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensortitlebarhideAPIRequest 隐藏titleBar API请求
// alibaba.interact.sensor.titlebarhide
//
// 隐藏titleBar
type AlibabainteractsensortitlebarhideAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensortitlebarhideRequest 初始化AlibabainteractsensortitlebarhideAPIRequest对象
func NewAlibabainteractsensortitlebarhideRequest() *AlibabainteractsensortitlebarhideAPIRequest {
	return &AlibabainteractsensortitlebarhideAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensortitlebarhideAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.titlebarhide"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensortitlebarhideAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensortitlebarhideAPIRequest) GetRawParams() model.Params {
	return r.Params
}
