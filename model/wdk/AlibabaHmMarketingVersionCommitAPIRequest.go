package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingVersionCommitAPIRequest 提交版本号 API请求
// alibaba.hm.marketing.version.commit
//
// 提交版本号，标识结束此版本操作
type AlibabaHmMarketingVersionCommitAPIRequest struct {
	model.Params
	// 版本号提交参数
	_param *SeasonVersionCommitParam
}

// NewAlibabaHmMarketingVersionCommitRequest 初始化AlibabaHmMarketingVersionCommitAPIRequest对象
func NewAlibabaHmMarketingVersionCommitRequest() *AlibabaHmMarketingVersionCommitAPIRequest {
	return &AlibabaHmMarketingVersionCommitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingVersionCommitAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingVersionCommitAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.version.commit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingVersionCommitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingVersionCommitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 版本号提交参数
func (r *AlibabaHmMarketingVersionCommitAPIRequest) SetParam(_param *SeasonVersionCommitParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingVersionCommitAPIRequest) GetParam() *SeasonVersionCommitParam {
	return r._param
}

var poolAlibabaHmMarketingVersionCommitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingVersionCommitRequest()
	},
}

// GetAlibabaHmMarketingVersionCommitRequest 从 sync.Pool 获取 AlibabaHmMarketingVersionCommitAPIRequest
func GetAlibabaHmMarketingVersionCommitAPIRequest() *AlibabaHmMarketingVersionCommitAPIRequest {
	return poolAlibabaHmMarketingVersionCommitAPIRequest.Get().(*AlibabaHmMarketingVersionCommitAPIRequest)
}

// ReleaseAlibabaHmMarketingVersionCommitAPIRequest 将 AlibabaHmMarketingVersionCommitAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingVersionCommitAPIRequest(v *AlibabaHmMarketingVersionCommitAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingVersionCommitAPIRequest.Put(v)
}
