package perfect

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaperfectperformanceitemqueryAPIRequest 商品完美履约信息查询 API请求
// alibaba.perfect.performance.item.query
//
// 同城零售商品完美履约信息查询
type AlibabaperfectperformanceitemqueryAPIRequest struct {
	model.Params
	// 查询入参
	_itemPerfectPerformanceQueryReq *ItemPerfectPerformanceQueryReq
}

// NewAlibabaperfectperformanceitemqueryRequest 初始化AlibabaperfectperformanceitemqueryAPIRequest对象
func NewAlibabaperfectperformanceitemqueryRequest() *AlibabaperfectperformanceitemqueryAPIRequest {
	return &AlibabaperfectperformanceitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaperfectperformanceitemqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.perfect.performance.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaperfectperformanceitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaperfectperformanceitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemPerfectPerformanceQueryReq is ItemPerfectPerformanceQueryReq Setter
// 查询入参
func (r *AlibabaperfectperformanceitemqueryAPIRequest) SetItemPerfectPerformanceQueryReq(_itemPerfectPerformanceQueryReq *ItemPerfectPerformanceQueryReq) error {
	r._itemPerfectPerformanceQueryReq = _itemPerfectPerformanceQueryReq
	r.Set("item_perfect_performance_query_req", _itemPerfectPerformanceQueryReq)
	return nil
}

// GetItemPerfectPerformanceQueryReq ItemPerfectPerformanceQueryReq Getter
func (r AlibabaperfectperformanceitemqueryAPIRequest) GetItemPerfectPerformanceQueryReq() *ItemPerfectPerformanceQueryReq {
	return r._itemPerfectPerformanceQueryReq
}
