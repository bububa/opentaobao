package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeSpecialUsersQueryAPIRequest 专属下单标记信息查询 API请求
// taobao.opentrade.special.users.query
//
// 专属下单标记信息查询
type TaobaoOpentradeSpecialUsersQueryAPIRequest struct {
	model.Params
	// 用户openId列表，多个以逗号(,)分割
	_openUserIds []string
	// 分页参数，每页大小
	_pageSize int64
	// 商品ID
	_itemId int64
	// 商品SKU ID，不存在传0
	_skuId int64
	// 用户状态
	_status string
	// 分页参数，当前页，以0开始
	_pageIndex int64
}

// NewTaobaoOpentradeSpecialUsersQueryRequest 初始化TaobaoOpentradeSpecialUsersQueryAPIRequest对象
func NewTaobaoOpentradeSpecialUsersQueryRequest() *TaobaoOpentradeSpecialUsersQueryAPIRequest {
	return &TaobaoOpentradeSpecialUsersQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.special.users.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OpenUserIds Setter
// 用户openId列表，多个以逗号(,)分割
func (r *TaobaoOpentradeSpecialUsersQueryAPIRequest) SetOpenUserIds(_openUserIds []string) error {
	r._openUserIds = _openUserIds
	r.Set("open_user_ids", _openUserIds)
	return nil
}

// Get OpenUserIds Getter
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetOpenUserIds() []string {
	return r._openUserIds
}

// Set is PageSize Setter
// 分页参数，每页大小
func (r *TaobaoOpentradeSpecialUsersQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is ItemId Setter
// 商品ID
func (r *TaobaoOpentradeSpecialUsersQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is SkuId Setter
// 商品SKU ID，不存在传0
func (r *TaobaoOpentradeSpecialUsersQueryAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// Get SkuId Getter
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// Set is Status Setter
// 用户状态
func (r *TaobaoOpentradeSpecialUsersQueryAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetStatus() string {
	return r._status
}

// Set is PageIndex Setter
// 分页参数，当前页，以0开始
func (r *TaobaoOpentradeSpecialUsersQueryAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// Get PageIndex Getter
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}
