package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapurcmallgoodssyncAPIRequest 第三方商家接入采购商城-商品同步 API请求
// alibaba.pur.cmall.goods.sync
//
// 第三方商家接入采购商城-商品同步
type AlibabapurcmallgoodssyncAPIRequest struct {
	model.Params
	// 产品对象
	_accessProductDto *AccessProductDto
	// 商品对象
	_accessGoodsDto *AccessGoodsDto
}

// NewAlibabapurcmallgoodssyncRequest 初始化AlibabapurcmallgoodssyncAPIRequest对象
func NewAlibabapurcmallgoodssyncRequest() *AlibabapurcmallgoodssyncAPIRequest {
	return &AlibabapurcmallgoodssyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapurcmallgoodssyncAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.cmall.goods.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapurcmallgoodssyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapurcmallgoodssyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessProductDto is AccessProductDto Setter
// 产品对象
func (r *AlibabapurcmallgoodssyncAPIRequest) SetAccessProductDto(_accessProductDto *AccessProductDto) error {
	r._accessProductDto = _accessProductDto
	r.Set("access_product_dto", _accessProductDto)
	return nil
}

// GetAccessProductDto AccessProductDto Getter
func (r AlibabapurcmallgoodssyncAPIRequest) GetAccessProductDto() *AccessProductDto {
	return r._accessProductDto
}

// SetAccessGoodsDto is AccessGoodsDto Setter
// 商品对象
func (r *AlibabapurcmallgoodssyncAPIRequest) SetAccessGoodsDto(_accessGoodsDto *AccessGoodsDto) error {
	r._accessGoodsDto = _accessGoodsDto
	r.Set("access_goods_dto", _accessGoodsDto)
	return nil
}

// GetAccessGoodsDto AccessGoodsDto Getter
func (r AlibabapurcmallgoodssyncAPIRequest) GetAccessGoodsDto() *AccessGoodsDto {
	return r._accessGoodsDto
}
