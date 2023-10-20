package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTraceUrlGetAPIRequest 溯源url透出 API请求
// alibaba.wdk.trace.url.get
//
// 根据shopId和skuCode返回商品溯源url
type AlibabaWdkTraceUrlGetAPIRequest struct {
	model.Params
	// barCode 或者skuCode
	_scanCode string
	// 所属门店code
	_shopId string
	// 来源编码
	_sourceCode string
}

// NewAlibabaWdkTraceUrlGetRequest 初始化AlibabaWdkTraceUrlGetAPIRequest对象
func NewAlibabaWdkTraceUrlGetRequest() *AlibabaWdkTraceUrlGetAPIRequest {
	return &AlibabaWdkTraceUrlGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkTraceUrlGetAPIRequest) Reset() {
	r._scanCode = ""
	r._shopId = ""
	r._sourceCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTraceUrlGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trace.url.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTraceUrlGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkTraceUrlGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScanCode is ScanCode Setter
// barCode 或者skuCode
func (r *AlibabaWdkTraceUrlGetAPIRequest) SetScanCode(_scanCode string) error {
	r._scanCode = _scanCode
	r.Set("scan_code", _scanCode)
	return nil
}

// GetScanCode ScanCode Getter
func (r AlibabaWdkTraceUrlGetAPIRequest) GetScanCode() string {
	return r._scanCode
}

// SetShopId is ShopId Setter
// 所属门店code
func (r *AlibabaWdkTraceUrlGetAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabaWdkTraceUrlGetAPIRequest) GetShopId() string {
	return r._shopId
}

// SetSourceCode is SourceCode Setter
// 来源编码
func (r *AlibabaWdkTraceUrlGetAPIRequest) SetSourceCode(_sourceCode string) error {
	r._sourceCode = _sourceCode
	r.Set("source_code", _sourceCode)
	return nil
}

// GetSourceCode SourceCode Getter
func (r AlibabaWdkTraceUrlGetAPIRequest) GetSourceCode() string {
	return r._sourceCode
}

var poolAlibabaWdkTraceUrlGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkTraceUrlGetRequest()
	},
}

// GetAlibabaWdkTraceUrlGetRequest 从 sync.Pool 获取 AlibabaWdkTraceUrlGetAPIRequest
func GetAlibabaWdkTraceUrlGetAPIRequest() *AlibabaWdkTraceUrlGetAPIRequest {
	return poolAlibabaWdkTraceUrlGetAPIRequest.Get().(*AlibabaWdkTraceUrlGetAPIRequest)
}

// ReleaseAlibabaWdkTraceUrlGetAPIRequest 将 AlibabaWdkTraceUrlGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkTraceUrlGetAPIRequest(v *AlibabaWdkTraceUrlGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkTraceUrlGetAPIRequest.Put(v)
}
