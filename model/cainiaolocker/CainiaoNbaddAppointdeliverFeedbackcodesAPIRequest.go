package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoNbaddAppointdeliverFeedbackcodesAPIRequest 服务质量反馈编码列表 API请求
// cainiao.nbadd.appointdeliver.feedbackcodes
//
// 服务质量反馈编码列表，建议获取数据后缓存在本地，定时刷新即可
type CainiaoNbaddAppointdeliverFeedbackcodesAPIRequest struct {
	model.Params
}

// NewCainiaoNbaddAppointdeliverFeedbackcodesRequest 初始化CainiaoNbaddAppointdeliverFeedbackcodesAPIRequest对象
func NewCainiaoNbaddAppointdeliverFeedbackcodesRequest() *CainiaoNbaddAppointdeliverFeedbackcodesAPIRequest {
	return &CainiaoNbaddAppointdeliverFeedbackcodesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoNbaddAppointdeliverFeedbackcodesAPIRequest) GetApiMethodName() string {
	return "cainiao.nbadd.appointdeliver.feedbackcodes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoNbaddAppointdeliverFeedbackcodesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoNbaddAppointdeliverFeedbackcodesAPIRequest) GetRawParams() model.Params {
	return r.Params
}
