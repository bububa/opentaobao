package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurCmallGoodsSyncAPIRequest 第三方商家接入采购商城-商品同步 API请求
// alibaba.pur.cmall.goods.sync
//
// 第三方商家接入采购商城-商品同步
type AlibabaPurCmallGoodsSyncAPIRequest struct {
	model.Params
	// 产品对象
	_accessProductDto *AccessProductDto
	// 商品对象
	_accessGoodsDto *AccessGoodsDto
}

// NewAlibabaPurCmallGoodsSyncRequest 初始化AlibabaPurCmallGoodsSyncAPIRequest对象
func NewAlibabaPurCmallGoodsSyncRequest() *AlibabaPurCmallGoodsSyncAPIRequest {
	return &AlibabaPurCmallGoodsSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPurCmallGoodsSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.cmall.goods.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPurCmallGoodsSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPurCmallGoodsSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessProductDto is AccessProductDto Setter
// 产品对象
func (r *AlibabaPurCmallGoodsSyncAPIRequest) SetAccessProductDto(_accessProductDto *AccessProductDto) error {
	r._accessProductDto = _accessProductDto
	r.Set("access_product_dto", _accessProductDto)
	return nil
}

// GetAccessProductDto AccessProductDto Getter
func (r AlibabaPurCmallGoodsSyncAPIRequest) GetAccessProductDto() *AccessProductDto {
	return r._accessProductDto
}

// SetAccessGoodsDto is AccessGoodsDto Setter
// 商品对象
func (r *AlibabaPurCmallGoodsSyncAPIRequest) SetAccessGoodsDto(_accessGoodsDto *AccessGoodsDto) error {
	r._accessGoodsDto = _accessGoodsDto
	r.Set("access_goods_dto", _accessGoodsDto)
	return nil
}

// GetAccessGoodsDto AccessGoodsDto Getter
func (r AlibabaPurCmallGoodsSyncAPIRequest) GetAccessGoodsDto() *AccessGoodsDto {
	return r._accessGoodsDto
}
