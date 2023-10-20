package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolActivityCreateAPIRequest 创建活动新接口 API请求
// alibaba.wdk.marketing.itempool.activity.create
//
// 创建活动新接口，支持新工具玩法
type AlibabaWdkMarketingItempoolActivityCreateAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemPoolActivity
}

// NewAlibabaWdkMarketingItempoolActivityCreateRequest 初始化AlibabaWdkMarketingItempoolActivityCreateAPIRequest对象
func NewAlibabaWdkMarketingItempoolActivityCreateRequest() *AlibabaWdkMarketingItempoolActivityCreateAPIRequest {
	return &AlibabaWdkMarketingItempoolActivityCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItempoolActivityCreateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolActivityCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolActivityCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItempoolActivityCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabaWdkMarketingItempoolActivityCreateAPIRequest) SetParam(_param *ItemPoolActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingItempoolActivityCreateAPIRequest) GetParam() *ItemPoolActivity {
	return r._param
}

var poolAlibabaWdkMarketingItempoolActivityCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItempoolActivityCreateRequest()
	},
}

// GetAlibabaWdkMarketingItempoolActivityCreateRequest 从 sync.Pool 获取 AlibabaWdkMarketingItempoolActivityCreateAPIRequest
func GetAlibabaWdkMarketingItempoolActivityCreateAPIRequest() *AlibabaWdkMarketingItempoolActivityCreateAPIRequest {
	return poolAlibabaWdkMarketingItempoolActivityCreateAPIRequest.Get().(*AlibabaWdkMarketingItempoolActivityCreateAPIRequest)
}

// ReleaseAlibabaWdkMarketingItempoolActivityCreateAPIRequest 将 AlibabaWdkMarketingItempoolActivityCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolActivityCreateAPIRequest(v *AlibabaWdkMarketingItempoolActivityCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolActivityCreateAPIRequest.Put(v)
}
