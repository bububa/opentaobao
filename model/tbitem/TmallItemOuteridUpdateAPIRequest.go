package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemouteridupdateAPIRequest 天猫商品/SKU商家编码更新接口 API请求
// tmall.item.outerid.update
//
// 天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名）
type TmallitemouteridupdateAPIRequest struct {
	model.Params
	// 商品SKU更新OuterId时候用的数据
	_skuOuters []UpdateSkuOuterId
	// 商品维度商家编码，如果不修改可以不传；清空请设置空串
	_outerId string
	// 商品ID
	_itemId int64
}

// NewTmallitemouteridupdateRequest 初始化TmallitemouteridupdateAPIRequest对象
func NewTmallitemouteridupdateRequest() *TmallitemouteridupdateAPIRequest {
	return &TmallitemouteridupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemouteridupdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.outerid.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemouteridupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemouteridupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuOuters is SkuOuters Setter
// 商品SKU更新OuterId时候用的数据
func (r *TmallitemouteridupdateAPIRequest) SetSkuOuters(_skuOuters []UpdateSkuOuterId) error {
	r._skuOuters = _skuOuters
	r.Set("sku_outers", _skuOuters)
	return nil
}

// GetSkuOuters SkuOuters Getter
func (r TmallitemouteridupdateAPIRequest) GetSkuOuters() []UpdateSkuOuterId {
	return r._skuOuters
}

// SetOuterId is OuterId Setter
// 商品维度商家编码，如果不修改可以不传；清空请设置空串
func (r *TmallitemouteridupdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TmallitemouteridupdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallitemouteridupdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemouteridupdateAPIRequest) GetItemId() int64 {
	return r._itemId
}
