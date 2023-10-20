package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoOrderQueryAPIRequest 查询订单基本信息 API请求
// alibaba.ele.fengniao.order.query
//
// 查询订单基本信息
type AlibabaEleFengniaoOrderQueryAPIRequest struct {
	model.Params
	// 参数
	_param *Param
}

// NewAlibabaEleFengniaoOrderQueryRequest 初始化AlibabaEleFengniaoOrderQueryAPIRequest对象
func NewAlibabaEleFengniaoOrderQueryRequest() *AlibabaEleFengniaoOrderQueryAPIRequest {
	return &AlibabaEleFengniaoOrderQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleFengniaoOrderQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleFengniaoOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 参数
func (r *AlibabaEleFengniaoOrderQueryAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaEleFengniaoOrderQueryAPIRequest) GetParam() *Param {
	return r._param
}

var poolAlibabaEleFengniaoOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleFengniaoOrderQueryRequest()
	},
}

// GetAlibabaEleFengniaoOrderQueryRequest 从 sync.Pool 获取 AlibabaEleFengniaoOrderQueryAPIRequest
func GetAlibabaEleFengniaoOrderQueryAPIRequest() *AlibabaEleFengniaoOrderQueryAPIRequest {
	return poolAlibabaEleFengniaoOrderQueryAPIRequest.Get().(*AlibabaEleFengniaoOrderQueryAPIRequest)
}

// ReleaseAlibabaEleFengniaoOrderQueryAPIRequest 将 AlibabaEleFengniaoOrderQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleFengniaoOrderQueryAPIRequest(v *AlibabaEleFengniaoOrderQueryAPIRequest) {
	v.Reset()
	poolAlibabaEleFengniaoOrderQueryAPIRequest.Put(v)
}
