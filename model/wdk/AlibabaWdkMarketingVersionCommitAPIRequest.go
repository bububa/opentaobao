package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingVersionCommitAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.version.commit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingVersionCommitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
