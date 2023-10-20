package wdkitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemMerchantskuQueryAPIRequest 商家商品信息查询 API请求
// alibaba.wdk.item.merchantsku.query
//
// 商家商品信息查询
type AlibabaWdkItemMerchantskuQueryAPIRequest struct {
	model.Params
	// 商品编码
	_skuCode string
	// 机构编码
	_orgCode string
}

// NewAlibabaWdkItemMerchantskuQueryRequest 初始化AlibabaWdkItemMerchantskuQueryAPIRequest对象
func NewAlibabaWdkItemMerchantskuQueryRequest() *AlibabaWdkItemMerchantskuQueryAPIRequest {
	return &AlibabaWdkItemMerchantskuQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkItemMerchantskuQueryAPIRequest) Reset() {
	r._skuCode = ""
	r._orgCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantskuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.merchantsku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantskuQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkItemMerchantskuQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemMerchantskuQueryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabaWdkItemMerchantskuQueryAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetOrgCode is OrgCode Setter
// 机构编码
func (r *AlibabaWdkItemMerchantskuQueryAPIRequest) SetOrgCode(_orgCode string) error {
	r._orgCode = _orgCode
	r.Set("org_code", _orgCode)
	return nil
}

// GetOrgCode OrgCode Getter
func (r AlibabaWdkItemMerchantskuQueryAPIRequest) GetOrgCode() string {
	return r._orgCode
}

var poolAlibabaWdkItemMerchantskuQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkItemMerchantskuQueryRequest()
	},
}

// GetAlibabaWdkItemMerchantskuQueryRequest 从 sync.Pool 获取 AlibabaWdkItemMerchantskuQueryAPIRequest
func GetAlibabaWdkItemMerchantskuQueryAPIRequest() *AlibabaWdkItemMerchantskuQueryAPIRequest {
	return poolAlibabaWdkItemMerchantskuQueryAPIRequest.Get().(*AlibabaWdkItemMerchantskuQueryAPIRequest)
}

// ReleaseAlibabaWdkItemMerchantskuQueryAPIRequest 将 AlibabaWdkItemMerchantskuQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkItemMerchantskuQueryAPIRequest(v *AlibabaWdkItemMerchantskuQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkItemMerchantskuQueryAPIRequest.Put(v)
}
