package perfect

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPerfectPerformanceItemQueryAPIRequest 商品完美履约信息查询 API请求
// alibaba.perfect.performance.item.query
//
// 同城零售商品完美履约信息查询
type AlibabaPerfectPerformanceItemQueryAPIRequest struct {
	model.Params
	// 查询入参
	_itemPerfectPerformanceQueryReq *ItemPerfectPerformanceQueryReq
}

// NewAlibabaPerfectPerformanceItemQueryRequest 初始化AlibabaPerfectPerformanceItemQueryAPIRequest对象
func NewAlibabaPerfectPerformanceItemQueryRequest() *AlibabaPerfectPerformanceItemQueryAPIRequest {
	return &AlibabaPerfectPerformanceItemQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPerfectPerformanceItemQueryAPIRequest) Reset() {
	r._itemPerfectPerformanceQueryReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPerfectPerformanceItemQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.perfect.performance.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPerfectPerformanceItemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPerfectPerformanceItemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemPerfectPerformanceQueryReq is ItemPerfectPerformanceQueryReq Setter
// 查询入参
func (r *AlibabaPerfectPerformanceItemQueryAPIRequest) SetItemPerfectPerformanceQueryReq(_itemPerfectPerformanceQueryReq *ItemPerfectPerformanceQueryReq) error {
	r._itemPerfectPerformanceQueryReq = _itemPerfectPerformanceQueryReq
	r.Set("item_perfect_performance_query_req", _itemPerfectPerformanceQueryReq)
	return nil
}

// GetItemPerfectPerformanceQueryReq ItemPerfectPerformanceQueryReq Getter
func (r AlibabaPerfectPerformanceItemQueryAPIRequest) GetItemPerfectPerformanceQueryReq() *ItemPerfectPerformanceQueryReq {
	return r._itemPerfectPerformanceQueryReq
}

var poolAlibabaPerfectPerformanceItemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPerfectPerformanceItemQueryRequest()
	},
}

// GetAlibabaPerfectPerformanceItemQueryRequest 从 sync.Pool 获取 AlibabaPerfectPerformanceItemQueryAPIRequest
func GetAlibabaPerfectPerformanceItemQueryAPIRequest() *AlibabaPerfectPerformanceItemQueryAPIRequest {
	return poolAlibabaPerfectPerformanceItemQueryAPIRequest.Get().(*AlibabaPerfectPerformanceItemQueryAPIRequest)
}

// ReleaseAlibabaPerfectPerformanceItemQueryAPIRequest 将 AlibabaPerfectPerformanceItemQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaPerfectPerformanceItemQueryAPIRequest(v *AlibabaPerfectPerformanceItemQueryAPIRequest) {
	v.Reset()
	poolAlibabaPerfectPerformanceItemQueryAPIRequest.Put(v)
}
