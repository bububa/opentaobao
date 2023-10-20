package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractisvgatewayAPIRequest isv调用gateway API请求
// alibaba.interact.isv.gateway
//
// isv能够调用jae本身的server
type AlibabainteractisvgatewayAPIRequest struct {
	model.Params
}

// NewAlibabainteractisvgatewayRequest 初始化AlibabainteractisvgatewayAPIRequest对象
func NewAlibabainteractisvgatewayRequest() *AlibabainteractisvgatewayAPIRequest {
	return &AlibabainteractisvgatewayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractisvgatewayAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.isv.gateway"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractisvgatewayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractisvgatewayAPIRequest) GetRawParams() model.Params {
	return r.Params
}
