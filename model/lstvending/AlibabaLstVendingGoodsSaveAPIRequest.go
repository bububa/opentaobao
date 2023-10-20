package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstvendinggoodssaveAPIRequest 自动售卖机商品回传 API请求
// alibaba.lst.vending.goods.save
//
// 零售通自动售卖机商品数据回流。
type AlibabalstvendinggoodssaveAPIRequest struct {
	model.Params
	// 商品信息
	_goodsDTOList []VendingGoodsDto
}

// NewAlibabalstvendinggoodssaveRequest 初始化AlibabalstvendinggoodssaveAPIRequest对象
func NewAlibabalstvendinggoodssaveRequest() *AlibabalstvendinggoodssaveAPIRequest {
	return &AlibabalstvendinggoodssaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstvendinggoodssaveAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.goods.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstvendinggoodssaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstvendinggoodssaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsDTOList is GoodsDTOList Setter
// 商品信息
func (r *AlibabalstvendinggoodssaveAPIRequest) SetGoodsDTOList(_goodsDTOList []VendingGoodsDto) error {
	r._goodsDTOList = _goodsDTOList
	r.Set("goods_d_t_o_list", _goodsDTOList)
	return nil
}

// GetGoodsDTOList GoodsDTOList Getter
func (r AlibabalstvendinggoodssaveAPIRequest) GetGoodsDTOList() []VendingGoodsDto {
	return r._goodsDTOList
}
