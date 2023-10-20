package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolCreateactivityAPIRequest 添加商品池活动 API请求
// alibaba.wdk.marketing.itempool.createactivity
//
// 添加商品池活动
type AlibabaWdkMarketingItempoolCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemPoolActivity
}

// NewAlibabaWdkMarketingItempoolCreateactivityRequest 初始化AlibabaWdkMarketingItempoolCreateactivityAPIRequest对象
func NewAlibabaWdkMarketingItempoolCreateactivityRequest() *AlibabaWdkMarketingItempoolCreateactivityAPIRequest {
	return &AlibabaWdkMarketingItempoolCreateactivityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItempoolCreateactivityAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolCreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolCreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItempoolCreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabaWdkMarketingItempoolCreateactivityAPIRequest) SetParam(_param *ItemPoolActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingItempoolCreateactivityAPIRequest) GetParam() *ItemPoolActivity {
	return r._param
}

var poolAlibabaWdkMarketingItempoolCreateactivityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItempoolCreateactivityRequest()
	},
}

// GetAlibabaWdkMarketingItempoolCreateactivityRequest 从 sync.Pool 获取 AlibabaWdkMarketingItempoolCreateactivityAPIRequest
func GetAlibabaWdkMarketingItempoolCreateactivityAPIRequest() *AlibabaWdkMarketingItempoolCreateactivityAPIRequest {
	return poolAlibabaWdkMarketingItempoolCreateactivityAPIRequest.Get().(*AlibabaWdkMarketingItempoolCreateactivityAPIRequest)
}

// ReleaseAlibabaWdkMarketingItempoolCreateactivityAPIRequest 将 AlibabaWdkMarketingItempoolCreateactivityAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolCreateactivityAPIRequest(v *AlibabaWdkMarketingItempoolCreateactivityAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolCreateactivityAPIRequest.Put(v)
}
