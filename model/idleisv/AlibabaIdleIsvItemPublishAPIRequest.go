package idleisv

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvItemPublishAPIRequest 服务商闲鱼商品发布 API请求
// alibaba.idle.isv.item.publish
//
// 服务商ISV闲鱼商品发布
type AlibabaIdleIsvItemPublishAPIRequest struct {
	model.Params
	// 商品数据参数
	_itemParam *IdleItemApiDo
}

// NewAlibabaIdleIsvItemPublishRequest 初始化AlibabaIdleIsvItemPublishAPIRequest对象
func NewAlibabaIdleIsvItemPublishRequest() *AlibabaIdleIsvItemPublishAPIRequest {
	return &AlibabaIdleIsvItemPublishAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleIsvItemPublishAPIRequest) Reset() {
	r._itemParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvItemPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.item.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvItemPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvItemPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemParam is ItemParam Setter
// 商品数据参数
func (r *AlibabaIdleIsvItemPublishAPIRequest) SetItemParam(_itemParam *IdleItemApiDo) error {
	r._itemParam = _itemParam
	r.Set("item_param", _itemParam)
	return nil
}

// GetItemParam ItemParam Getter
func (r AlibabaIdleIsvItemPublishAPIRequest) GetItemParam() *IdleItemApiDo {
	return r._itemParam
}

var poolAlibabaIdleIsvItemPublishAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleIsvItemPublishRequest()
	},
}

// GetAlibabaIdleIsvItemPublishRequest 从 sync.Pool 获取 AlibabaIdleIsvItemPublishAPIRequest
func GetAlibabaIdleIsvItemPublishAPIRequest() *AlibabaIdleIsvItemPublishAPIRequest {
	return poolAlibabaIdleIsvItemPublishAPIRequest.Get().(*AlibabaIdleIsvItemPublishAPIRequest)
}

// ReleaseAlibabaIdleIsvItemPublishAPIRequest 将 AlibabaIdleIsvItemPublishAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleIsvItemPublishAPIRequest(v *AlibabaIdleIsvItemPublishAPIRequest) {
	v.Reset()
	poolAlibabaIdleIsvItemPublishAPIRequest.Put(v)
}
