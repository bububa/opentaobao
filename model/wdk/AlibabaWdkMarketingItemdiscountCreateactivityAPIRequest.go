package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest 创建商品特价活动 API请求
// alibaba.wdk.marketing.itemdiscount.createactivity
//
// 创建商品特价活动
type AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemDiscountActivityRequest
}

// NewAlibabaWdkMarketingItemdiscountCreateactivityRequest 初始化AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest对象
func NewAlibabaWdkMarketingItemdiscountCreateactivityRequest() *AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest {
	return &AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itemdiscount.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest) SetParam(_param *ItemDiscountActivityRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest) GetParam() *ItemDiscountActivityRequest {
	return r._param
}

var poolAlibabaWdkMarketingItemdiscountCreateactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItemdiscountCreateactivityRequest()
	},
}

// GetAlibabaWdkMarketingItemdiscountCreateactivityRequest 从 sync.Pool 获取 AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest
func GetAlibabaWdkMarketingItemdiscountCreateactivityAPIRequest() *AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest {
	return poolAlibabaWdkMarketingItemdiscountCreateactivityAPIRequest.Get().(*AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest)
}

// ReleaseAlibabaWdkMarketingItemdiscountCreateactivityAPIRequest 将 AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItemdiscountCreateactivityAPIRequest(v *AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItemdiscountCreateactivityAPIRequest.Put(v)
}
