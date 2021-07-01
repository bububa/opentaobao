package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeGroupActivtiyQueryAPIRequest
组团活动信息查询 API请求
taobao.opentrade.group.activtiy.query

组团购场景下，团购活动信息查询 */
type TaobaoOpentradeGroupActivtiyQueryAPIRequest struct {
	model.Params
	// 分页参数，每页大小
	_pageSize int64
	// 商品ID
	_itemId int64
	// 分页参数，当前页，以0开始
	_pageIndex int64
	// 组团活动id
	_groupActivityId int64
}

// NewTaobaoOpentradeGroupActivtiyQueryRequest 初始化TaobaoOpentradeGroupActivtiyQueryAPIRequest对象
func NewTaobaoOpentradeGroupActivtiyQueryRequest() *TaobaoOpentradeGroupActivtiyQueryAPIRequest {
	return &TaobaoOpentradeGroupActivtiyQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupActivtiyQueryAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.group.activtiy.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupActivtiyQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PageSize Setter
// 分页参数，每页大小
func (r *TaobaoOpentradeGroupActivtiyQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoOpentradeGroupActivtiyQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is ItemId Setter
// 商品ID
func (r *TaobaoOpentradeGroupActivtiyQueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoOpentradeGroupActivtiyQueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is PageIndex Setter
// 分页参数，当前页，以0开始
func (r *TaobaoOpentradeGroupActivtiyQueryAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// Get PageIndex Getter
func (r TaobaoOpentradeGroupActivtiyQueryAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// Set is GroupActivityId Setter
// 组团活动id
func (r *TaobaoOpentradeGroupActivtiyQueryAPIRequest) SetGroupActivityId(_groupActivityId int64) error {
	r._groupActivityId = _groupActivityId
	r.Set("group_activity_id", _groupActivityId)
	return nil
}

// Get GroupActivityId Getter
func (r TaobaoOpentradeGroupActivtiyQueryAPIRequest) GetGroupActivityId() int64 {
	return r._groupActivityId
}
