package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsSkuGetAPIRequest 商品信息查询 API请求
// taobao.wlb.wms.sku.get
//
// 商品信息查询
type TaobaoWlbWmsSkuGetAPIRequest struct {
	model.Params
	// 菜鸟商品ID,与itemcode必须有一个值不为空
	_itemCode string
	// 商家商品编码,与itemid必须有一个值不为空
	_itemId string
	// 货主ID
	_ownerUserId string
}

// NewTaobaoWlbWmsSkuGetRequest 初始化TaobaoWlbWmsSkuGetAPIRequest对象
func NewTaobaoWlbWmsSkuGetRequest() *TaobaoWlbWmsSkuGetAPIRequest {
	return &TaobaoWlbWmsSkuGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsSkuGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.sku.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsSkuGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemCode is ItemCode Setter
// 菜鸟商品ID,与itemcode必须有一个值不为空
func (r *TaobaoWlbWmsSkuGetAPIRequest) SetItemCode(_itemCode string) error {
	r._itemCode = _itemCode
	r.Set("item_code", _itemCode)
	return nil
}

// GetItemCode ItemCode Getter
func (r TaobaoWlbWmsSkuGetAPIRequest) GetItemCode() string {
	return r._itemCode
}

// SetItemId is ItemId Setter
// 商家商品编码,与itemid必须有一个值不为空
func (r *TaobaoWlbWmsSkuGetAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoWlbWmsSkuGetAPIRequest) GetItemId() string {
	return r._itemId
}

// SetOwnerUserId is OwnerUserId Setter
// 货主ID
func (r *TaobaoWlbWmsSkuGetAPIRequest) SetOwnerUserId(_ownerUserId string) error {
	r._ownerUserId = _ownerUserId
	r.Set("owner_user_id", _ownerUserId)
	return nil
}

// GetOwnerUserId OwnerUserId Getter
func (r TaobaoWlbWmsSkuGetAPIRequest) GetOwnerUserId() string {
	return r._ownerUserId
}
