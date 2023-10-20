package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingVersionCommitAPIRequest 提交版本号 API请求
// alibaba.wdk.marketing.version.commit
//
// 提交版本号，标识结束此版本操作
type AlibabaWdkMarketingVersionCommitAPIRequest struct {
	model.Params
	// 版本号提交参数
	_param *SeasonVersionCommitParam
}

// NewAlibabaWdkMarketingVersionCommitRequest 初始化AlibabaWdkMarketingVersionCommitAPIRequest对象
func NewAlibabaWdkMarketingVersionCommitRequest() *AlibabaWdkMarketingVersionCommitAPIRequest {
	return &AlibabaWdkMarketingVersionCommitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingVersionCommitAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingVersionCommitAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.version.commit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingVersionCommitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingVersionCommitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 版本号提交参数
func (r *AlibabaWdkMarketingVersionCommitAPIRequest) SetParam(_param *SeasonVersionCommitParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingVersionCommitAPIRequest) GetParam() *SeasonVersionCommitParam {
	return r._param
}

var poolAlibabaWdkMarketingVersionCommitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingVersionCommitRequest()
	},
}

// GetAlibabaWdkMarketingVersionCommitRequest 从 sync.Pool 获取 AlibabaWdkMarketingVersionCommitAPIRequest
func GetAlibabaWdkMarketingVersionCommitAPIRequest() *AlibabaWdkMarketingVersionCommitAPIRequest {
	return poolAlibabaWdkMarketingVersionCommitAPIRequest.Get().(*AlibabaWdkMarketingVersionCommitAPIRequest)
}

// ReleaseAlibabaWdkMarketingVersionCommitAPIRequest 将 AlibabaWdkMarketingVersionCommitAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingVersionCommitAPIRequest(v *AlibabaWdkMarketingVersionCommitAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingVersionCommitAPIRequest.Put(v)
}
