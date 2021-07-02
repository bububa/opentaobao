package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemMapQueryAPIRequest 查找IC商品或分销商品与后端商品的关联信息 API请求
// taobao.scitem.map.query
//
// 查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku
type TaobaoScitemMapQueryAPIRequest struct {
	model.Params
	// map_type为1：前台ic商品id<br/>map_type为2：分销productid
	_itemId int64
	// map_type为1：前台ic商品skuid <br/>map_type为2：分销商品的skuid
	_skuId int64
}

// NewTaobaoScitemMapQueryRequest 初始化TaobaoScitemMapQueryAPIRequest对象
func NewTaobaoScitemMapQueryRequest() *TaobaoScitemMapQueryAPIRequest {
	return &TaobaoScitemMapQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoScitemMapQueryAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.map.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoScitemMapQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemId is ItemId Setter
// map_type为1：前台ic商品id<br/>map_type为2：分销productid
func (r *TaobaoScitemMapQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoScitemMapQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// map_type为1：前台ic商品skuid <br/>map_type为2：分销商品的skuid
func (r *TaobaoScitemMapQueryAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoScitemMapQueryAPIRequest) GetSkuId() int64 {
	return r._skuId
}
