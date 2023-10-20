package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemItemGetAPIRequest 获取全渠道门店商品 API请求
// taobao.omniitem.item.get
//
// 通过门店id/类目id/商品id单个或多个参数组合查询全渠道门店商品信息
type TaobaoOmniitemItemGetAPIRequest struct {
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

// NewTaobaoOmniitemItemGetRequest 初始化TaobaoOmniitemItemGetAPIRequest对象
func NewTaobaoOmniitemItemGetRequest() *TaobaoOmniitemItemGetAPIRequest {
	return &TaobaoOmniitemItemGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniitemItemGetAPIRequest) Reset() {
	r._outerId = ""
	r._pageNo = 0
	r._pageSize = 0
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemItemGetAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.item.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemItemGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniitemItemGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 可选，指定获取的商品外部id
func (r *TaobaoOmniitemItemGetAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoOmniitemItemGetAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetPageNo is PageNo Setter
// 分页当前页数
func (r *TaobaoOmniitemItemGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoOmniitemItemGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页单页大小
func (r *TaobaoOmniitemItemGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoOmniitemItemGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetItemId is ItemId Setter
// 可选，指定获取的商品id
func (r *TaobaoOmniitemItemGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoOmniitemItemGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoOmniitemItemGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniitemItemGetRequest()
	},
}

// GetTaobaoOmniitemItemGetRequest 从 sync.Pool 获取 TaobaoOmniitemItemGetAPIRequest
func GetTaobaoOmniitemItemGetAPIRequest() *TaobaoOmniitemItemGetAPIRequest {
	return poolTaobaoOmniitemItemGetAPIRequest.Get().(*TaobaoOmniitemItemGetAPIRequest)
}

// ReleaseTaobaoOmniitemItemGetAPIRequest 将 TaobaoOmniitemItemGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniitemItemGetAPIRequest(v *TaobaoOmniitemItemGetAPIRequest) {
	v.Reset()
	poolTaobaoOmniitemItemGetAPIRequest.Put(v)
}
