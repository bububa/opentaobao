package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingversiongenerateAPIRequest 生成发布使用的版本号 API请求
// alibaba.hm.marketing.version.generate
//
// 生成发布使用的版本号
type AlibabahmmarketingversiongenerateAPIRequest struct {
	model.Params
	// 档期版本号参数信息
	_param *SeasonVersionParam
}

// NewAlibabahmmarketingversiongenerateRequest 初始化AlibabahmmarketingversiongenerateAPIRequest对象
func NewAlibabahmmarketingversiongenerateRequest() *AlibabahmmarketingversiongenerateAPIRequest {
	return &AlibabahmmarketingversiongenerateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingversiongenerateAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.version.generate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingversiongenerateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingversiongenerateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 档期版本号参数信息
func (r *AlibabahmmarketingversiongenerateAPIRequest) SetParam(_param *SeasonVersionParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingversiongenerateAPIRequest) GetParam() *SeasonVersionParam {
	return r._param
}
