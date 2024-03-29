package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest 删除商品特价活动 API请求
// alibaba.wdk.marketing.itemdiscount.deleteactivity
//
// 删除商品特价活动
type AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityRequest
}

// NewAlibabaWdkMarketingItemdiscountDeleteactivityRequest 初始化AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest对象
func NewAlibabaWdkMarketingItemdiscountDeleteactivityRequest() *AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest {
	return &AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itemdiscount.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest) SetParam(_param *CommonActivityRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest) GetParam() *CommonActivityRequest {
	return r._param
}

var poolAlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItemdiscountDeleteactivityRequest()
	},
}

// GetAlibabaWdkMarketingItemdiscountDeleteactivityRequest 从 sync.Pool 获取 AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest
func GetAlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest() *AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest {
	return poolAlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest.Get().(*AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest)
}

// ReleaseAlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest 将 AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest(v *AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest.Put(v)
}
