package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuMerchantskuScrollQueryAPIRequest 商家商品批量查询接口 API请求
// alibaba.wdk.sku.merchantsku.scroll.query
//
// 提供主档商品数据接口查询
type AlibabaWdkSkuMerchantskuScrollQueryAPIRequest struct {
	model.Params
	// HM
	_orgNo string
	// 第一次为null，后面从结果中获取
	_scrollId string
}

// NewAlibabaWdkSkuMerchantskuScrollQueryRequest 初始化AlibabaWdkSkuMerchantskuScrollQueryAPIRequest对象
func NewAlibabaWdkSkuMerchantskuScrollQueryRequest() *AlibabaWdkSkuMerchantskuScrollQueryAPIRequest {
	return &AlibabaWdkSkuMerchantskuScrollQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.merchantsku.scroll.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrgNo Setter
// HM
func (r *AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) SetOrgNo(_orgNo string) error {
	r._orgNo = _orgNo
	r.Set("org_no", _orgNo)
	return nil
}

// Get OrgNo Getter
func (r AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) GetOrgNo() string {
	return r._orgNo
}

// Set is ScrollId Setter
// 第一次为null，后面从结果中获取
func (r *AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) SetScrollId(_scrollId string) error {
	r._scrollId = _scrollId
	r.Set("scroll_id", _scrollId)
	return nil
}

// Get ScrollId Getter
func (r AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) GetScrollId() string {
	return r._scrollId
}
