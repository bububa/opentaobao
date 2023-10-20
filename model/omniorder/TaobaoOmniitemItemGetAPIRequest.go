package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniitemitemgetAPIRequest 获取全渠道门店商品 API请求
// taobao.omniitem.item.get
//
// 通过门店id/类目id/商品id单个或多个参数组合查询全渠道门店商品信息
type TaobaoomniitemitemgetAPIRequest struct {
	model.Params
	// 可选，指定获取的商品外部id
	_outerId string
	// 分页当前页数
	_pageNo int64
	// 分页单页大小
	_pageSize int64
	// 可选，指定获取的商品id
	_itemId int64
}

// NewTaobaoomniitemitemgetRequest 初始化TaobaoomniitemitemgetAPIRequest对象
func NewTaobaoomniitemitemgetRequest() *TaobaoomniitemitemgetAPIRequest {
	return &TaobaoomniitemitemgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniitemitemgetAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniitemitemgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniitemitemgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 可选，指定获取的商品外部id
func (r *TaobaoomniitemitemgetAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoomniitemitemgetAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetPageNo is PageNo Setter
// 分页当前页数
func (r *TaobaoomniitemitemgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoomniitemitemgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页单页大小
func (r *TaobaoomniitemitemgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoomniitemitemgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetItemId is ItemId Setter
// 可选，指定获取的商品id
func (r *TaobaoomniitemitemgetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoomniitemitemgetAPIRequest) GetItemId() int64 {
	return r._itemId
}
