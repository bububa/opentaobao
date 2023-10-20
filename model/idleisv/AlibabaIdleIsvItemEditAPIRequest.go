package idleisv

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvItemEditAPIRequest 服务商闲鱼商品编辑 API请求
// alibaba.idle.isv.item.edit
//
// 服务商ISV闲鱼商品编辑操作
type AlibabaIdleIsvItemEditAPIRequest struct {
	model.Params
	// 商品数据参数
	_param *IdleItemApiDo
}

// NewAlibabaIdleIsvItemEditRequest 初始化AlibabaIdleIsvItemEditAPIRequest对象
func NewAlibabaIdleIsvItemEditRequest() *AlibabaIdleIsvItemEditAPIRequest {
	return &AlibabaIdleIsvItemEditAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleIsvItemEditAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvItemEditAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.item.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvItemEditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvItemEditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 商品数据参数
func (r *AlibabaIdleIsvItemEditAPIRequest) SetParam(_param *IdleItemApiDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaIdleIsvItemEditAPIRequest) GetParam() *IdleItemApiDo {
	return r._param
}

var poolAlibabaIdleIsvItemEditAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleIsvItemEditRequest()
	},
}

// GetAlibabaIdleIsvItemEditRequest 从 sync.Pool 获取 AlibabaIdleIsvItemEditAPIRequest
func GetAlibabaIdleIsvItemEditAPIRequest() *AlibabaIdleIsvItemEditAPIRequest {
	return poolAlibabaIdleIsvItemEditAPIRequest.Get().(*AlibabaIdleIsvItemEditAPIRequest)
}

// ReleaseAlibabaIdleIsvItemEditAPIRequest 将 AlibabaIdleIsvItemEditAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleIsvItemEditAPIRequest(v *AlibabaIdleIsvItemEditAPIRequest) {
	v.Reset()
	poolAlibabaIdleIsvItemEditAPIRequest.Put(v)
}
