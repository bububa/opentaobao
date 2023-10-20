package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTransferpayQueryAPIRequest 闲鱼转账结果查询 API请求
// alibaba.idle.transferpay.query
//
// 商家业务 转账支付的结果查询
type AlibabaIdleTransferpayQueryAPIRequest struct {
	model.Params
	// 入参
	_param *PayQueryRequest
}

// NewAlibabaIdleTransferpayQueryRequest 初始化AlibabaIdleTransferpayQueryAPIRequest对象
func NewAlibabaIdleTransferpayQueryRequest() *AlibabaIdleTransferpayQueryAPIRequest {
	return &AlibabaIdleTransferpayQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleTransferpayQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTransferpayQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.transferpay.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTransferpayQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleTransferpayQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaIdleTransferpayQueryAPIRequest) SetParam(_param *PayQueryRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaIdleTransferpayQueryAPIRequest) GetParam() *PayQueryRequest {
	return r._param
}

var poolAlibabaIdleTransferpayQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleTransferpayQueryRequest()
	},
}

// GetAlibabaIdleTransferpayQueryRequest 从 sync.Pool 获取 AlibabaIdleTransferpayQueryAPIRequest
func GetAlibabaIdleTransferpayQueryAPIRequest() *AlibabaIdleTransferpayQueryAPIRequest {
	return poolAlibabaIdleTransferpayQueryAPIRequest.Get().(*AlibabaIdleTransferpayQueryAPIRequest)
}

// ReleaseAlibabaIdleTransferpayQueryAPIRequest 将 AlibabaIdleTransferpayQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleTransferpayQueryAPIRequest(v *AlibabaIdleTransferpayQueryAPIRequest) {
	v.Reset()
	poolAlibabaIdleTransferpayQueryAPIRequest.Put(v)
}
