package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvitempublishAPIRequest 服务商闲鱼商品发布 API请求
// alibaba.idle.isv.item.publish
//
// 服务商ISV闲鱼商品发布
type AlibabaidleisvitempublishAPIRequest struct {
	model.Params
	// 商品数据参数
	_itemParam *IdleItemApiDo
}

// NewAlibabaidleisvitempublishRequest 初始化AlibabaidleisvitempublishAPIRequest对象
func NewAlibabaidleisvitempublishRequest() *AlibabaidleisvitempublishAPIRequest {
	return &AlibabaidleisvitempublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvitempublishAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.item.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvitempublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvitempublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemParam is ItemParam Setter
// 商品数据参数
func (r *AlibabaidleisvitempublishAPIRequest) SetItemParam(_itemParam *IdleItemApiDo) error {
	r._itemParam = _itemParam
	r.Set("item_param", _itemParam)
	return nil
}

// GetItemParam ItemParam Getter
func (r AlibabaidleisvitempublishAPIRequest) GetItemParam() *IdleItemApiDo {
	return r._itemParam
}
