package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelelementssearchAPIRequest 商家元素搜索 API请求
// alitrip.travel.elements.search
//
// 提供商家维护的景点、酒店、餐饮等元素搜索
type AlitriptravelelementssearchAPIRequest struct {
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

// NewAlitriptravelelementssearchRequest 初始化AlitriptravelelementssearchAPIRequest对象
func NewAlitriptravelelementssearchRequest() *AlitriptravelelementssearchAPIRequest {
	return &AlitriptravelelementssearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptravelelementssearchAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.elements.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptravelelementssearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptravelelementssearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询关键词
func (r *AlitriptravelelementssearchAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlitriptravelelementssearchAPIRequest) GetQuery() string {
	return r._query
}

// SetSellerId is SellerId Setter
// 商家id
func (r *AlitriptravelelementssearchAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r AlitriptravelelementssearchAPIRequest) GetSellerId() int64 {
	return r._sellerId
}

// SetCount is Count Setter
// 查询数量，限制100
func (r *AlitriptravelelementssearchAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// GetCount Count Getter
func (r AlitriptravelelementssearchAPIRequest) GetCount() int64 {
	return r._count
}

// SetType is Type Setter
// 资源类型
func (r *AlitriptravelelementssearchAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlitriptravelelementssearchAPIRequest) GetType() int64 {
	return r._type
}
