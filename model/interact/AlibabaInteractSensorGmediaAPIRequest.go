package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorgmediaAPIRequest gmedia API请求
// alibaba.interact.sensor.gmedia
//
// 媒体功能
type AlibabainteractsensorgmediaAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorgmediaRequest 初始化AlibabainteractsensorgmediaAPIRequest对象
func NewAlibabainteractsensorgmediaRequest() *AlibabainteractsensorgmediaAPIRequest {
	return &AlibabainteractsensorgmediaAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorgmediaAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.gmedia"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorgmediaAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorgmediaAPIRequest) GetRawParams() model.Params {
	return r.Params
}
