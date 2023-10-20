package logistic

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleFengniaoOrderPushAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
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

var poolAlibabaEleFengniaoOrderPushAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleFengniaoOrderPushRequest()
	},
}

// GetAlibabaEleFengniaoOrderPushRequest 从 sync.Pool 获取 AlibabaEleFengniaoOrderPushAPIRequest
func GetAlibabaEleFengniaoOrderPushAPIRequest() *AlibabaEleFengniaoOrderPushAPIRequest {
	return poolAlibabaEleFengniaoOrderPushAPIRequest.Get().(*AlibabaEleFengniaoOrderPushAPIRequest)
}

// ReleaseAlibabaEleFengniaoOrderPushAPIRequest 将 AlibabaEleFengniaoOrderPushAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleFengniaoOrderPushAPIRequest(v *AlibabaEleFengniaoOrderPushAPIRequest) {
	v.Reset()
	poolAlibabaEleFengniaoOrderPushAPIRequest.Put(v)
}
