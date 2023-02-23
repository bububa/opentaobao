package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingVersionGenerateAPIRequest 生成发布使用的版本号 API请求
// alibaba.wdk.marketing.version.generate
//
// 生成发布使用的版本号
type AlibabaWdkMarketingVersionGenerateAPIRequest struct {
	model.Params
	// 档期版本号参数信息
	_param *SeasonVersionParam
}

// NewAlibabaWdkMarketingVersionGenerateRequest 初始化AlibabaWdkMarketingVersionGenerateAPIRequest对象
func NewAlibabaWdkMarketingVersionGenerateRequest() *AlibabaWdkMarketingVersionGenerateAPIRequest {
	return &AlibabaWdkMarketingVersionGenerateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingVersionGenerateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.version.generate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingVersionGenerateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingVersionGenerateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 档期版本号参数信息
func (r *AlibabaWdkMarketingVersionGenerateAPIRequest) SetParam(_param *SeasonVersionParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingVersionGenerateAPIRequest) GetParam() *SeasonVersionParam {
	return r._param
}
