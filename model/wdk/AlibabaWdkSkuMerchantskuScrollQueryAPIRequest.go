package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) Reset() {
	r._orgNo = ""
	r._scrollId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.merchantsku.scroll.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrgNo is OrgNo Setter
// HM
func (r *AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) SetOrgNo(_orgNo string) error {
	r._orgNo = _orgNo
	r.Set("org_no", _orgNo)
	return nil
}

// GetOrgNo OrgNo Getter
func (r AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) GetOrgNo() string {
	return r._orgNo
}

// SetScrollId is ScrollId Setter
// 第一次为null，后面从结果中获取
func (r *AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) SetScrollId(_scrollId string) error {
	r._scrollId = _scrollId
	r.Set("scroll_id", _scrollId)
	return nil
}

// GetScrollId ScrollId Getter
func (r AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) GetScrollId() string {
	return r._scrollId
}

var poolAlibabaWdkSkuMerchantskuScrollQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuMerchantskuScrollQueryRequest()
	},
}

// GetAlibabaWdkSkuMerchantskuScrollQueryRequest 从 sync.Pool 获取 AlibabaWdkSkuMerchantskuScrollQueryAPIRequest
func GetAlibabaWdkSkuMerchantskuScrollQueryAPIRequest() *AlibabaWdkSkuMerchantskuScrollQueryAPIRequest {
	return poolAlibabaWdkSkuMerchantskuScrollQueryAPIRequest.Get().(*AlibabaWdkSkuMerchantskuScrollQueryAPIRequest)
}

// ReleaseAlibabaWdkSkuMerchantskuScrollQueryAPIRequest 将 AlibabaWdkSkuMerchantskuScrollQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuMerchantskuScrollQueryAPIRequest(v *AlibabaWdkSkuMerchantskuScrollQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuMerchantskuScrollQueryAPIRequest.Put(v)
}
