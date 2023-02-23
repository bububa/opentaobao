package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractWindvaneCallAPIRequest Weex页面调用windvane API请求
// alibaba.interact.windvane.call
//
// 客户端鉴权使用，实际不会发送或接收数据
type AlibabaInteractWindvaneCallAPIRequest struct {
	model.Params
	// 客户端鉴权使用，实际不会发送或接收数据
	_unNamed string
}

// NewAlibabaInteractWindvaneCallRequest 初始化AlibabaInteractWindvaneCallAPIRequest对象
func NewAlibabaInteractWindvaneCallRequest() *AlibabaInteractWindvaneCallAPIRequest {
	return &AlibabaInteractWindvaneCallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractWindvaneCallAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.windvane.call"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractWindvaneCallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractWindvaneCallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUnNamed is UnNamed Setter
// 客户端鉴权使用，实际不会发送或接收数据
func (r *AlibabaInteractWindvaneCallAPIRequest) SetUnNamed(_unNamed string) error {
	r._unNamed = _unNamed
	r.Set("un_named", _unNamed)
	return nil
}

// GetUnNamed UnNamed Getter
func (r AlibabaInteractWindvaneCallAPIRequest) GetUnNamed() string {
	return r._unNamed
}
