package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingversiongenerateAPIRequest 生成发布使用的版本号 API请求
// alibaba.wdk.marketing.version.generate
//
// 生成发布使用的版本号
type AlibabawdkmarketingversiongenerateAPIRequest struct {
	model.Params
	// 档期版本号参数信息
	_param *SeasonVersionParam
}

// NewAlibabawdkmarketingversiongenerateRequest 初始化AlibabawdkmarketingversiongenerateAPIRequest对象
func NewAlibabawdkmarketingversiongenerateRequest() *AlibabawdkmarketingversiongenerateAPIRequest {
	return &AlibabawdkmarketingversiongenerateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingversiongenerateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.version.generate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingversiongenerateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingversiongenerateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 档期版本号参数信息
func (r *AlibabawdkmarketingversiongenerateAPIRequest) SetParam(_param *SeasonVersionParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingversiongenerateAPIRequest) GetParam() *SeasonVersionParam {
	return r._param
}
