package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoOrderPushAPIRequest 推送订单 API请求
// alibaba.ele.fengniao.order.push
//
// 推送淘宝订单至蜂鸟开放平台配送
type AlibabaEleFengniaoOrderPushAPIRequest struct {
	model.Params
	// 参数param
	_param *Param
}

// NewAlibabaEleFengniaoOrderPushRequest 初始化AlibabaEleFengniaoOrderPushAPIRequest对象
func NewAlibabaEleFengniaoOrderPushRequest() *AlibabaEleFengniaoOrderPushAPIRequest {
	return &AlibabaEleFengniaoOrderPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoOrderPushAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.order.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoOrderPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleFengniaoOrderPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 参数param
func (r *AlibabaEleFengniaoOrderPushAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaEleFengniaoOrderPushAPIRequest) GetParam() *Param {
	return r._param
}
