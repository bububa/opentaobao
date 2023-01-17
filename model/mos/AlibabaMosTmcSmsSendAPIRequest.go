package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosTmcSmsSendAPIRequest 银泰TMC发送短信 API请求
// alibaba.mos.tmc.sms.send
//
// 银泰TMC发送短信
type AlibabaMosTmcSmsSendAPIRequest struct {
	model.Params
	// 入参
	_param0 *SmsSendMessageDto
}

// NewAlibabaMosTmcSmsSendRequest 初始化AlibabaMosTmcSmsSendAPIRequest对象
func NewAlibabaMosTmcSmsSendRequest() *AlibabaMosTmcSmsSendAPIRequest {
	return &AlibabaMosTmcSmsSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosTmcSmsSendAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.tmc.sms.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosTmcSmsSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosTmcSmsSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *AlibabaMosTmcSmsSendAPIRequest) SetParam0(_param0 *SmsSendMessageDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaMosTmcSmsSendAPIRequest) GetParam0() *SmsSendMessageDto {
	return r._param0
}
