package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest 创建买赠活动 API请求
// alibaba.wdk.marketing.itembuygift.createactivity
//
// 创建买赠活动
type AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemBuyGiftActivity
}

// NewAlibabaWdkMarketingItembuygiftCreateactivityRequest 初始化AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest对象
func NewAlibabaWdkMarketingItembuygiftCreateactivityRequest() *AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest {
	return &AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itembuygift.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest) SetParam(_param *ItemBuyGiftActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest) GetParam() *ItemBuyGiftActivity {
	return r._param
}

var poolAlibabaWdkMarketingItembuygiftCreateactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItembuygiftCreateactivityRequest()
	},
}

// GetAlibabaWdkMarketingItembuygiftCreateactivityRequest 从 sync.Pool 获取 AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest
func GetAlibabaWdkMarketingItembuygiftCreateactivityAPIRequest() *AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest {
	return poolAlibabaWdkMarketingItembuygiftCreateactivityAPIRequest.Get().(*AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest)
}

// ReleaseAlibabaWdkMarketingItembuygiftCreateactivityAPIRequest 将 AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItembuygiftCreateactivityAPIRequest(v *AlibabaWdkMarketingItembuygiftCreateactivityAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItembuygiftCreateactivityAPIRequest.Put(v)
}
