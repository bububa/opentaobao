package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensortoastAPIRequest toast API请求
// alibaba.interact.sensor.toast
//
// toast提示
type AlibabainteractsensortoastAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensortoastRequest 初始化AlibabainteractsensortoastAPIRequest对象
func NewAlibabainteractsensortoastRequest() *AlibabainteractsensortoastAPIRequest {
	return &AlibabainteractsensortoastAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensortoastAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.toast"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensortoastAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensortoastAPIRequest) GetRawParams() model.Params {
	return r.Params
}
