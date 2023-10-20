package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingversioncommitAPIRequest 提交版本号 API请求
// alibaba.hm.marketing.version.commit
//
// 提交版本号，标识结束此版本操作
type AlibabahmmarketingversioncommitAPIRequest struct {
	model.Params
	// 版本号提交参数
	_param *SeasonVersionCommitParam
}

// NewAlibabahmmarketingversioncommitRequest 初始化AlibabahmmarketingversioncommitAPIRequest对象
func NewAlibabahmmarketingversioncommitRequest() *AlibabahmmarketingversioncommitAPIRequest {
	return &AlibabahmmarketingversioncommitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingversioncommitAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.version.commit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingversioncommitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingversioncommitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 版本号提交参数
func (r *AlibabahmmarketingversioncommitAPIRequest) SetParam(_param *SeasonVersionCommitParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingversioncommitAPIRequest) GetParam() *SeasonVersionCommitParam {
	return r._param
}
