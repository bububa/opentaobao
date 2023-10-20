package opentrade

import (
	"net/url"
	"sync"

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
	// 用户状态
	_status string
	// 分页参数，每页大小
	_pageSize int64
	// 商品ID
	_itemId int64
	// 商品SKU ID，不存在传0
	_skuId int64
	// 分页参数，当前页，以0开始
	_pageIndex int64
}

// NewTaobaoOpentradeSpecialUsersQueryRequest 初始化TaobaoOpentradeSpecialUsersQueryAPIRequest对象
func NewTaobaoOpentradeSpecialUsersQueryRequest() *TaobaoOpentradeSpecialUsersQueryAPIRequest {
	return &TaobaoOpentradeSpecialUsersQueryAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpentradeSpecialUsersQueryAPIRequest) Reset() {
	r._openUserIds = r._openUserIds[:0]
	r._status = ""
	r._pageSize = 0
	r._itemId = 0
	r._skuId = 0
	r._pageIndex = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.special.users.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenUserIds is OpenUserIds Setter
// 用户openId列表，多个以逗号(,)分割
func (r *TaobaoOpentradeSpecialUsersQueryAPIRequest) SetOpenUserIds(_openUserIds []string) error {
	r._openUserIds = _openUserIds
	r.Set("open_user_ids", _openUserIds)
	return nil
}

// GetOpenUserIds OpenUserIds Getter
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetOpenUserIds() []string {
	return r._openUserIds
}

// SetStatus is Status Setter
// 用户状态
func (r *TaobaoOpentradeSpecialUsersQueryAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetStatus() string {
	return r._status
}

// SetPageSize is PageSize Setter
// 分页参数，每页大小
func (r *TaobaoOpentradeSpecialUsersQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoOpentradeSpecialUsersQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// 商品SKU ID，不存在传0
func (r *TaobaoOpentradeSpecialUsersQueryAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetPageIndex is PageIndex Setter
// 分页参数，当前页，以0开始
func (r *TaobaoOpentradeSpecialUsersQueryAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoOpentradeSpecialUsersQueryAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

var poolTaobaoOpentradeSpecialUsersQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpentradeSpecialUsersQueryRequest()
	},
}

// GetTaobaoOpentradeSpecialUsersQueryRequest 从 sync.Pool 获取 TaobaoOpentradeSpecialUsersQueryAPIRequest
func GetTaobaoOpentradeSpecialUsersQueryAPIRequest() *TaobaoOpentradeSpecialUsersQueryAPIRequest {
	return poolTaobaoOpentradeSpecialUsersQueryAPIRequest.Get().(*TaobaoOpentradeSpecialUsersQueryAPIRequest)
}

// ReleaseTaobaoOpentradeSpecialUsersQueryAPIRequest 将 TaobaoOpentradeSpecialUsersQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpentradeSpecialUsersQueryAPIRequest(v *TaobaoOpentradeSpecialUsersQueryAPIRequest) {
	v.Reset()
	poolTaobaoOpentradeSpecialUsersQueryAPIRequest.Put(v)
}
