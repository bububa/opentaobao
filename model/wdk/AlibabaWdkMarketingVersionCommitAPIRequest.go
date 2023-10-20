package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingversioncommitAPIRequest 提交版本号 API请求
// alibaba.wdk.marketing.version.commit
//
// 提交版本号，标识结束此版本操作
type AlibabawdkmarketingversioncommitAPIRequest struct {
	model.Params
	// 版本号提交参数
	_param *SeasonVersionCommitParam
}

// NewAlibabawdkmarketingversioncommitRequest 初始化AlibabawdkmarketingversioncommitAPIRequest对象
func NewAlibabawdkmarketingversioncommitRequest() *AlibabawdkmarketingversioncommitAPIRequest {
	return &AlibabawdkmarketingversioncommitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingversioncommitAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.version.commit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingversioncommitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingversioncommitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 版本号提交参数
func (r *AlibabawdkmarketingversioncommitAPIRequest) SetParam(_param *SeasonVersionCommitParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingversioncommitAPIRequest) GetParam() *SeasonVersionCommitParam {
	return r._param
}
