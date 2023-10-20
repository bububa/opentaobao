package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaonbaddappointdeliverfeedbackcodesAPIRequest 服务质量反馈编码列表 API请求
// cainiao.nbadd.appointdeliver.feedbackcodes
//
// 服务质量反馈编码列表，建议获取数据后缓存在本地，定时刷新即可
type CainiaonbaddappointdeliverfeedbackcodesAPIRequest struct {
	model.Params
}

// NewCainiaonbaddappointdeliverfeedbackcodesRequest 初始化CainiaonbaddappointdeliverfeedbackcodesAPIRequest对象
func NewCainiaonbaddappointdeliverfeedbackcodesRequest() *CainiaonbaddappointdeliverfeedbackcodesAPIRequest {
	return &CainiaonbaddappointdeliverfeedbackcodesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaonbaddappointdeliverfeedbackcodesAPIRequest) GetApiMethodName() string {
	return "cainiao.nbadd.appointdeliver.feedbackcodes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaonbaddappointdeliverfeedbackcodesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaonbaddappointdeliverfeedbackcodesAPIRequest) GetRawParams() model.Params {
	return r.Params
}
