package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest 查找特价活动 API请求
// alibaba.wdk.marketing.itemdiscount.queryactivity
//
// 查找特价活动
type AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *CommonActivityParam
}

// NewAlibabaWdkMarketingItemdiscountQueryactivityRequest 初始化AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest对象
func NewAlibabaWdkMarketingItemdiscountQueryactivityRequest() *AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest {
	return &AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itemdiscount.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}

var poolAlibabaWdkMarketingItemdiscountQueryactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItemdiscountQueryactivityRequest()
	},
}

// GetAlibabaWdkMarketingItemdiscountQueryactivityRequest 从 sync.Pool 获取 AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest
func GetAlibabaWdkMarketingItemdiscountQueryactivityAPIRequest() *AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest {
	return poolAlibabaWdkMarketingItemdiscountQueryactivityAPIRequest.Get().(*AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest)
}

// ReleaseAlibabaWdkMarketingItemdiscountQueryactivityAPIRequest 将 AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItemdiscountQueryactivityAPIRequest(v *AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItemdiscountQueryactivityAPIRequest.Put(v)
}
