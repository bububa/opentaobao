package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoscitemmapqueryAPIRequest 查找IC商品或分销商品与后端商品的关联信息 API请求
// taobao.scitem.map.query
//
// 查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku
type TaobaoscitemmapqueryAPIRequest struct {
	model.Params
	// map_type为1：前台ic商品id<br/>map_type为2：分销productid
	_itemId int64
	// map_type为1：前台ic商品skuid <br/>map_type为2：分销商品的skuid
	_skuId int64
}

// NewTaobaoscitemmapqueryRequest 初始化TaobaoscitemmapqueryAPIRequest对象
func NewTaobaoscitemmapqueryRequest() *TaobaoscitemmapqueryAPIRequest {
	return &TaobaoscitemmapqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoscitemmapqueryAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.map.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoscitemmapqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoscitemmapqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// map_type为1：前台ic商品id&lt;br/&gt;map_type为2：分销productid
func (r *TaobaoscitemmapqueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoscitemmapqueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// map_type为1：前台ic商品skuid &lt;br/&gt;map_type为2：分销商品的skuid
func (r *TaobaoscitemmapqueryAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoscitemmapqueryAPIRequest) GetSkuId() int64 {
	return r._skuId
}
