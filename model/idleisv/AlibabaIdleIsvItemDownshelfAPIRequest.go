package idleisv

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvItemDownshelfAPIRequest 服务商闲鱼商品下架 API请求
// alibaba.idle.isv.item.downshelf
//
// 供外部服务商ISV进行闲鱼商品下架操作
type AlibabaIdleIsvItemDownshelfAPIRequest struct {
	model.Params
	// 返回结果result
	_param *IdleItemBaseApiDo
}

// NewAlibabaIdleIsvItemDownshelfRequest 初始化AlibabaIdleIsvItemDownshelfAPIRequest对象
func NewAlibabaIdleIsvItemDownshelfRequest() *AlibabaIdleIsvItemDownshelfAPIRequest {
	return &AlibabaIdleIsvItemDownshelfAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleIsvItemDownshelfAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvItemDownshelfAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.item.downshelf"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvItemDownshelfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvItemDownshelfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 返回结果result
func (r *AlibabaIdleIsvItemDownshelfAPIRequest) SetParam(_param *IdleItemBaseApiDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaIdleIsvItemDownshelfAPIRequest) GetParam() *IdleItemBaseApiDo {
	return r._param
}

var poolAlibabaIdleIsvItemDownshelfAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleIsvItemDownshelfRequest()
	},
}

// GetAlibabaIdleIsvItemDownshelfRequest 从 sync.Pool 获取 AlibabaIdleIsvItemDownshelfAPIRequest
func GetAlibabaIdleIsvItemDownshelfAPIRequest() *AlibabaIdleIsvItemDownshelfAPIRequest {
	return poolAlibabaIdleIsvItemDownshelfAPIRequest.Get().(*AlibabaIdleIsvItemDownshelfAPIRequest)
}

// ReleaseAlibabaIdleIsvItemDownshelfAPIRequest 将 AlibabaIdleIsvItemDownshelfAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleIsvItemDownshelfAPIRequest(v *AlibabaIdleIsvItemDownshelfAPIRequest) {
	v.Reset()
	poolAlibabaIdleIsvItemDownshelfAPIRequest.Put(v)
}
