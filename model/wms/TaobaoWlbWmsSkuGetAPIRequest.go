package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwmsskugetAPIRequest 商品信息查询 API请求
// taobao.wlb.wms.sku.get
//
// 商品信息查询
type TaobaowlbwmsskugetAPIRequest struct {
	model.Params
	// 菜鸟商品ID,与itemcode必须有一个值不为空
	_itemCode string
	// 商家商品编码,与itemid必须有一个值不为空
	_itemId string
	// 货主ID
	_ownerUserId string
}

// NewTaobaowlbwmsskugetRequest 初始化TaobaowlbwmsskugetAPIRequest对象
func NewTaobaowlbwmsskugetRequest() *TaobaowlbwmsskugetAPIRequest {
	return &TaobaowlbwmsskugetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwmsskugetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.sku.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwmsskugetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwmsskugetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemCode is ItemCode Setter
// 菜鸟商品ID,与itemcode必须有一个值不为空
func (r *TaobaowlbwmsskugetAPIRequest) SetItemCode(_itemCode string) error {
	r._itemCode = _itemCode
	r.Set("item_code", _itemCode)
	return nil
}

// GetItemCode ItemCode Getter
func (r TaobaowlbwmsskugetAPIRequest) GetItemCode() string {
	return r._itemCode
}

// SetItemId is ItemId Setter
// 商家商品编码,与itemid必须有一个值不为空
func (r *TaobaowlbwmsskugetAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaowlbwmsskugetAPIRequest) GetItemId() string {
	return r._itemId
}

// SetOwnerUserId is OwnerUserId Setter
// 货主ID
func (r *TaobaowlbwmsskugetAPIRequest) SetOwnerUserId(_ownerUserId string) error {
	r._ownerUserId = _ownerUserId
	r.Set("owner_user_id", _ownerUserId)
	return nil
}

// GetOwnerUserId OwnerUserId Getter
func (r TaobaowlbwmsskugetAPIRequest) GetOwnerUserId() string {
	return r._ownerUserId
}
