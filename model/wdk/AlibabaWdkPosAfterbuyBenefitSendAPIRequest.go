package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkposafterbuybenefitsendAPIRequest 生态pos购后发放权益 API请求
// alibaba.wdk.pos.afterbuy.benefit.send
//
// 生态pos购后发放权益接口开放
type AlibabawdkposafterbuybenefitsendAPIRequest struct {
	model.Params
	// 入参
	_sendBenefitParam *IsvSendBenefitParam
}

// NewAlibabawdkposafterbuybenefitsendRequest 初始化AlibabawdkposafterbuybenefitsendAPIRequest对象
func NewAlibabawdkposafterbuybenefitsendRequest() *AlibabawdkposafterbuybenefitsendAPIRequest {
	return &AlibabawdkposafterbuybenefitsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkposafterbuybenefitsendAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.afterbuy.benefit.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkposafterbuybenefitsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkposafterbuybenefitsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSendBenefitParam is SendBenefitParam Setter
// 入参
func (r *AlibabawdkposafterbuybenefitsendAPIRequest) SetSendBenefitParam(_sendBenefitParam *IsvSendBenefitParam) error {
	r._sendBenefitParam = _sendBenefitParam
	r.Set("send_benefit_param", _sendBenefitParam)
	return nil
}

// GetSendBenefitParam SendBenefitParam Getter
func (r AlibabawdkposafterbuybenefitsendAPIRequest) GetSendBenefitParam() *IsvSendBenefitParam {
	return r._sendBenefitParam
}
