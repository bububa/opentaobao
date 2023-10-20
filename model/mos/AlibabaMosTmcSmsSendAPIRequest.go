package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamostmcsmssendAPIRequest 银泰TMC发送短信 API请求
// alibaba.mos.tmc.sms.send
//
// 银泰TMC发送短信
type AlibabamostmcsmssendAPIRequest struct {
	model.Params
	// 入参
	_param0 *SmsSendMessageDto
}

// NewAlibabamostmcsmssendRequest 初始化AlibabamostmcsmssendAPIRequest对象
func NewAlibabamostmcsmssendRequest() *AlibabamostmcsmssendAPIRequest {
	return &AlibabamostmcsmssendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamostmcsmssendAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.tmc.sms.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamostmcsmssendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamostmcsmssendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *AlibabamostmcsmssendAPIRequest) SetParam0(_param0 *SmsSendMessageDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabamostmcsmssendAPIRequest) GetParam0() *SmsSendMessageDto {
	return r._param0
}
