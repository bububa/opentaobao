package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemTraceUrlGetAPIRequest 根据shopId和skuCode返回商品静态溯源url API请求
// alibaba.wdk.item.trace.url.get
//
// 根据shopId和skuCode返回商品静态溯源url
type AlibabaWdkItemTraceUrlGetAPIRequest struct {
	model.Params
	// 商品skuCode
	_skuCode string
	// 所属门店code
	_shopId string
	// 来源编码
	_sourceCode string
}

// NewAlibabaWdkItemTraceUrlGetRequest 初始化AlibabaWdkItemTraceUrlGetAPIRequest对象
func NewAlibabaWdkItemTraceUrlGetRequest() *AlibabaWdkItemTraceUrlGetAPIRequest {
	return &AlibabaWdkItemTraceUrlGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkItemTraceUrlGetAPIRequest) Reset() {
	r._skuCode = ""
	r._shopId = ""
	r._sourceCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemTraceUrlGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.trace.url.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemTraceUrlGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkItemTraceUrlGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCode is SkuCode Setter
// 商品skuCode
func (r *AlibabaWdkItemTraceUrlGetAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabaWdkItemTraceUrlGetAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetShopId is ShopId Setter
// 所属门店code
func (r *AlibabaWdkItemTraceUrlGetAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabaWdkItemTraceUrlGetAPIRequest) GetShopId() string {
	return r._shopId
}

// SetSourceCode is SourceCode Setter
// 来源编码
func (r *AlibabaWdkItemTraceUrlGetAPIRequest) SetSourceCode(_sourceCode string) error {
	r._sourceCode = _sourceCode
	r.Set("source_code", _sourceCode)
	return nil
}

// GetSourceCode SourceCode Getter
func (r AlibabaWdkItemTraceUrlGetAPIRequest) GetSourceCode() string {
	return r._sourceCode
}

var poolAlibabaWdkItemTraceUrlGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkItemTraceUrlGetRequest()
	},
}

// GetAlibabaWdkItemTraceUrlGetRequest 从 sync.Pool 获取 AlibabaWdkItemTraceUrlGetAPIRequest
func GetAlibabaWdkItemTraceUrlGetAPIRequest() *AlibabaWdkItemTraceUrlGetAPIRequest {
	return poolAlibabaWdkItemTraceUrlGetAPIRequest.Get().(*AlibabaWdkItemTraceUrlGetAPIRequest)
}

// ReleaseAlibabaWdkItemTraceUrlGetAPIRequest 将 AlibabaWdkItemTraceUrlGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkItemTraceUrlGetAPIRequest(v *AlibabaWdkItemTraceUrlGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkItemTraceUrlGetAPIRequest.Put(v)
}
