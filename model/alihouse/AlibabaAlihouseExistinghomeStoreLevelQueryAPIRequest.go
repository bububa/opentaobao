package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomestorelevelqueryAPIRequest 门店等级评分查询 API请求
// alibaba.alihouse.existinghome.store.level.query
//
// 门店等级评分查询
type AlibabaalihouseexistinghomestorelevelqueryAPIRequest struct {
	model.Params
	// 外部门店id列表，不能超过200个
	_outerStoreIds string
	// 行业属性  2-二手房  3-租房
	_businessType int64
}

// NewAlibabaalihouseexistinghomestorelevelqueryRequest 初始化AlibabaalihouseexistinghomestorelevelqueryAPIRequest对象
func NewAlibabaalihouseexistinghomestorelevelqueryRequest() *AlibabaalihouseexistinghomestorelevelqueryAPIRequest {
	return &AlibabaalihouseexistinghomestorelevelqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomestorelevelqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.store.level.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomestorelevelqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomestorelevelqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterStoreIds is OuterStoreIds Setter
// 外部门店id列表，不能超过200个
func (r *AlibabaalihouseexistinghomestorelevelqueryAPIRequest) SetOuterStoreIds(_outerStoreIds string) error {
	r._outerStoreIds = _outerStoreIds
	r.Set("outer_store_ids", _outerStoreIds)
	return nil
}

// GetOuterStoreIds OuterStoreIds Getter
func (r AlibabaalihouseexistinghomestorelevelqueryAPIRequest) GetOuterStoreIds() string {
	return r._outerStoreIds
}

// SetBusinessType is BusinessType Setter
// 行业属性  2-二手房  3-租房
func (r *AlibabaalihouseexistinghomestorelevelqueryAPIRequest) SetBusinessType(_businessType int64) error {
	r._businessType = _businessType
	r.Set("business_type", _businessType)
	return nil
}

// GetBusinessType BusinessType Getter
func (r AlibabaalihouseexistinghomestorelevelqueryAPIRequest) GetBusinessType() int64 {
	return r._businessType
}
