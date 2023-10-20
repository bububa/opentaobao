package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemmerchantskuqueryAPIRequest 商家商品信息查询 API请求
// alibaba.wdk.item.merchantsku.query
//
// 商家商品信息查询
type AlibabawdkitemmerchantskuqueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 机构编码
	_orgCode string
}

// NewAlibabawdkitemmerchantskuqueryRequest 初始化AlibabawdkitemmerchantskuqueryAPIRequest对象
func NewAlibabawdkitemmerchantskuqueryRequest() *AlibabawdkitemmerchantskuqueryAPIRequest {
	return &AlibabawdkitemmerchantskuqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemmerchantskuqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.merchantsku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemmerchantskuqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemmerchantskuqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabawdkitemmerchantskuqueryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabawdkitemmerchantskuqueryAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetOrgCode is OrgCode Setter
// 机构编码
func (r *AlibabawdkitemmerchantskuqueryAPIRequest) SetOrgCode(_orgCode string) error {
	r._orgCode = _orgCode
	r.Set("org_code", _orgCode)
	return nil
}

// GetOrgCode OrgCode Getter
func (r AlibabawdkitemmerchantskuqueryAPIRequest) GetOrgCode() string {
	return r._orgCode
}
