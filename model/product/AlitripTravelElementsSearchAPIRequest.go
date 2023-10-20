package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelElementsSearchAPIRequest 商家元素搜索 API请求
// alitrip.travel.elements.search
//
// 提供商家维护的景点、酒店、餐饮等元素搜索
type AlitripTravelElementsSearchAPIRequest struct {
	model.Params
	// 查询关键词
	_query string
	// 商家id
	_sellerId int64
	// 查询数量，限制100
	_count int64
	// 资源类型
	_type int64
}

// NewAlitripTravelElementsSearchRequest 初始化AlitripTravelElementsSearchAPIRequest对象
func NewAlitripTravelElementsSearchRequest() *AlitripTravelElementsSearchAPIRequest {
	return &AlitripTravelElementsSearchAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTravelElementsSearchAPIRequest) Reset() {
	r._query = ""
	r._sellerId = 0
	r._count = 0
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelElementsSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.elements.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelElementsSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelElementsSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询关键词
func (r *AlitripTravelElementsSearchAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlitripTravelElementsSearchAPIRequest) GetQuery() string {
	return r._query
}

// SetSellerId is SellerId Setter
// 商家id
func (r *AlitripTravelElementsSearchAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r AlitripTravelElementsSearchAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// SetCount is Count Setter
// 查询数量，限制100
func (r *AlitripTravelElementsSearchAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// GetCount Count Getter
func (r AlitripTravelElementsSearchAPIRequest) GetCount() int64 {
	return r._count
}

// SetType is Type Setter
// 资源类型
func (r *AlitripTravelElementsSearchAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlitripTravelElementsSearchAPIRequest) GetType() int64 {
	return r._type
}

var poolAlitripTravelElementsSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTravelElementsSearchRequest()
	},
}

// GetAlitripTravelElementsSearchRequest 从 sync.Pool 获取 AlitripTravelElementsSearchAPIRequest
func GetAlitripTravelElementsSearchAPIRequest() *AlitripTravelElementsSearchAPIRequest {
	return poolAlitripTravelElementsSearchAPIRequest.Get().(*AlitripTravelElementsSearchAPIRequest)
}

// ReleaseAlitripTravelElementsSearchAPIRequest 将 AlitripTravelElementsSearchAPIRequest 放入 sync.Pool
func ReleaseAlitripTravelElementsSearchAPIRequest(v *AlitripTravelElementsSearchAPIRequest) {
	v.Reset()
	poolAlitripTravelElementsSearchAPIRequest.Put(v)
}
