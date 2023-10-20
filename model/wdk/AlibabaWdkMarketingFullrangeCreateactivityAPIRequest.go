package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeCreateactivityAPIRequest 创建全场活动 API请求
// alibaba.wdk.marketing.fullrange.createactivity
//
// 创建全场活动
type AlibabaWdkMarketingFullrangeCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *FullRangeActivity
}

// NewAlibabaWdkMarketingFullrangeCreateactivityRequest 初始化AlibabaWdkMarketingFullrangeCreateactivityAPIRequest对象
func NewAlibabaWdkMarketingFullrangeCreateactivityRequest() *AlibabaWdkMarketingFullrangeCreateactivityAPIRequest {
	return &AlibabaWdkMarketingFullrangeCreateactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingFullrangeCreateactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeCreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.fullrange.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeCreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingFullrangeCreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabaWdkMarketingFullrangeCreateactivityAPIRequest) SetParam(_param *FullRangeActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingFullrangeCreateactivityAPIRequest) GetParam() *FullRangeActivity {
	return r._param
}

var poolAlibabaWdkMarketingFullrangeCreateactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingFullrangeCreateactivityRequest()
	},
}

// GetAlibabaWdkMarketingFullrangeCreateactivityRequest 从 sync.Pool 获取 AlibabaWdkMarketingFullrangeCreateactivityAPIRequest
func GetAlibabaWdkMarketingFullrangeCreateactivityAPIRequest() *AlibabaWdkMarketingFullrangeCreateactivityAPIRequest {
	return poolAlibabaWdkMarketingFullrangeCreateactivityAPIRequest.Get().(*AlibabaWdkMarketingFullrangeCreateactivityAPIRequest)
}

// ReleaseAlibabaWdkMarketingFullrangeCreateactivityAPIRequest 将 AlibabaWdkMarketingFullrangeCreateactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingFullrangeCreateactivityAPIRequest(v *AlibabaWdkMarketingFullrangeCreateactivityAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingFullrangeCreateactivityAPIRequest.Put(v)
}
