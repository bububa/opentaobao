package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIRequest 天猫服务平台商家投诉单问题列表查询 API请求
// tmall.servicecenter.anomalyrecourse.homedecoration.questioncode.query
//
// 天猫服务平台商家投诉单问题列表查询
type TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIRequest struct {
	model.Params
	// 服务code
	_serviceCode string
}

// NewTmallservicecenteranomalyrecoursehomedecorationquestioncodequeryRequest 初始化TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIRequest对象
func NewTmallservicecenteranomalyrecoursehomedecorationquestioncodequeryRequest() *TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIRequest {
	return &TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.homedecoration.questioncode.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceCode is ServiceCode Setter
// 服务code
func (r *TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r TmallservicecenteranomalyrecoursehomedecorationquestioncodequeryAPIRequest) GetServiceCode() string {
	return r._serviceCode
}
