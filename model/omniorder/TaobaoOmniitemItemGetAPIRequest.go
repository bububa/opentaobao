package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemItemGetAPIRequest 获取全渠道门店商品 API请求
// taobao.omniitem.item.get
//
// 通过门店id/类目id/商品id单个或多个参数组合查询全渠道门店商品信息
type TaobaoOmniitemItemGetAPIRequest struct {
	model.Params
	// 分页当前页数
	_pageNo int64
	// 分页单页大小
	_pageSize int64
	// 可选，指定获取的商品id
	_itemId int64
	// 可选，指定获取的商品外部id
	_outerId string
}

// NewTaobaoOmniitemItemGetRequest 初始化TaobaoOmniitemItemGetAPIRequest对象
func NewTaobaoOmniitemItemGetRequest() *TaobaoOmniitemItemGetAPIRequest {
	return &TaobaoOmniitemItemGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemItemGetAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemItemGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PageNo Setter
// 分页当前页数
func (r *TaobaoOmniitemItemGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoOmniitemItemGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 分页单页大小
func (r *TaobaoOmniitemItemGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoOmniitemItemGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is ItemId Setter
// 可选，指定获取的商品id
func (r *TaobaoOmniitemItemGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoOmniitemItemGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is OuterId Setter
// 可选，指定获取的商品外部id
func (r *TaobaoOmniitemItemGetAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r TaobaoOmniitemItemGetAPIRequest) GetOuterId() string {
	return r._outerId
}
