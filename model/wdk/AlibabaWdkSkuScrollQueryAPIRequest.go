package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuScrollQueryAPIRequest 门店商品批量游标方式查询接口 API请求
// alibaba.wdk.sku.scroll.query
//
// 通过游标方式批量获取门店商品信息，包括商品条码，商品名称，价格，会员价等信息。
type AlibabaWdkSkuScrollQueryAPIRequest struct {
	model.Params
	// 商家类目编码
	_merchantCatCode string
	// 门店编码
	_ouCode string
	// 游标：第一次请求不用填写，否则请填写上一次请求返回的值，直到获取到足够的数据
	_scrollId string
	// 英文逗号分隔的商品编码，最多20个。如果配合门店字段使用，直接非游标方式返回商品数据
	_skuCodes string
}

// NewAlibabaWdkSkuScrollQueryRequest 初始化AlibabaWdkSkuScrollQueryAPIRequest对象
func NewAlibabaWdkSkuScrollQueryRequest() *AlibabaWdkSkuScrollQueryAPIRequest {
	return &AlibabaWdkSkuScrollQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuScrollQueryAPIRequest) Reset() {
	r._merchantCatCode = ""
	r._ouCode = ""
	r._scrollId = ""
	r._skuCodes = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuScrollQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.scroll.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuScrollQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuScrollQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMerchantCatCode is MerchantCatCode Setter
// 商家类目编码
func (r *AlibabaWdkSkuScrollQueryAPIRequest) SetMerchantCatCode(_merchantCatCode string) error {
	r._merchantCatCode = _merchantCatCode
	r.Set("merchant_cat_code", _merchantCatCode)
	return nil
}

// GetMerchantCatCode MerchantCatCode Getter
func (r AlibabaWdkSkuScrollQueryAPIRequest) GetMerchantCatCode() string {
	return r._merchantCatCode
}

// SetOuCode is OuCode Setter
// 门店编码
func (r *AlibabaWdkSkuScrollQueryAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r AlibabaWdkSkuScrollQueryAPIRequest) GetOuCode() string {
	return r._ouCode
}

// SetScrollId is ScrollId Setter
// 游标：第一次请求不用填写，否则请填写上一次请求返回的值，直到获取到足够的数据
func (r *AlibabaWdkSkuScrollQueryAPIRequest) SetScrollId(_scrollId string) error {
	r._scrollId = _scrollId
	r.Set("scroll_id", _scrollId)
	return nil
}

// GetScrollId ScrollId Getter
func (r AlibabaWdkSkuScrollQueryAPIRequest) GetScrollId() string {
	return r._scrollId
}

// SetSkuCodes is SkuCodes Setter
// 英文逗号分隔的商品编码，最多20个。如果配合门店字段使用，直接非游标方式返回商品数据
func (r *AlibabaWdkSkuScrollQueryAPIRequest) SetSkuCodes(_skuCodes string) error {
	r._skuCodes = _skuCodes
	r.Set("sku_codes", _skuCodes)
	return nil
}

// GetSkuCodes SkuCodes Getter
func (r AlibabaWdkSkuScrollQueryAPIRequest) GetSkuCodes() string {
	return r._skuCodes
}

var poolAlibabaWdkSkuScrollQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuScrollQueryRequest()
	},
}

// GetAlibabaWdkSkuScrollQueryRequest 从 sync.Pool 获取 AlibabaWdkSkuScrollQueryAPIRequest
func GetAlibabaWdkSkuScrollQueryAPIRequest() *AlibabaWdkSkuScrollQueryAPIRequest {
	return poolAlibabaWdkSkuScrollQueryAPIRequest.Get().(*AlibabaWdkSkuScrollQueryAPIRequest)
}

// ReleaseAlibabaWdkSkuScrollQueryAPIRequest 将 AlibabaWdkSkuScrollQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuScrollQueryAPIRequest(v *AlibabaWdkSkuScrollQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuScrollQueryAPIRequest.Put(v)
}
