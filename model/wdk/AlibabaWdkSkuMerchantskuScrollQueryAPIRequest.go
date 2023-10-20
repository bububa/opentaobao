package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskumerchantskuscrollqueryAPIRequest 商家商品批量查询接口 API请求
// alibaba.wdk.sku.merchantsku.scroll.query
//
// 提供主档商品数据接口查询
type AlibabawdkskumerchantskuscrollqueryAPIRequest struct {
	model.Params
	// HM
	_orgNo string
	// 第一次为null，后面从结果中获取
	_scrollId string
}

// NewAlibabawdkskumerchantskuscrollqueryRequest 初始化AlibabawdkskumerchantskuscrollqueryAPIRequest对象
func NewAlibabawdkskumerchantskuscrollqueryRequest() *AlibabawdkskumerchantskuscrollqueryAPIRequest {
	return &AlibabawdkskumerchantskuscrollqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskumerchantskuscrollqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.merchantsku.scroll.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskumerchantskuscrollqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskumerchantskuscrollqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrgNo is OrgNo Setter
// HM
func (r *AlibabawdkskumerchantskuscrollqueryAPIRequest) SetOrgNo(_orgNo string) error {
	r._orgNo = _orgNo
	r.Set("org_no", _orgNo)
	return nil
}

// GetOrgNo OrgNo Getter
func (r AlibabawdkskumerchantskuscrollqueryAPIRequest) GetOrgNo() string {
	return r._orgNo
}

// SetScrollId is ScrollId Setter
// 第一次为null，后面从结果中获取
func (r *AlibabawdkskumerchantskuscrollqueryAPIRequest) SetScrollId(_scrollId string) error {
	r._scrollId = _scrollId
	r.Set("scroll_id", _scrollId)
	return nil
}

// GetScrollId ScrollId Getter
func (r AlibabawdkskumerchantskuscrollqueryAPIRequest) GetScrollId() string {
	return r._scrollId
}
