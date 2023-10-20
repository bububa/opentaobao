package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectchannelphoneAPIRequest 新房渠道电话数据同步 API请求
// alibaba.alihouse.newhome.project.channelphone
//
// 新房渠道电话数据同步
type AlibabaalihousenewhomeprojectchannelphoneAPIRequest struct {
	model.Params
	// 渠道电话对象
	_channelPhoneDto *ProjectChannelPhoneDto
}

// NewAlibabaalihousenewhomeprojectchannelphoneRequest 初始化AlibabaalihousenewhomeprojectchannelphoneAPIRequest对象
func NewAlibabaalihousenewhomeprojectchannelphoneRequest() *AlibabaalihousenewhomeprojectchannelphoneAPIRequest {
	return &AlibabaalihousenewhomeprojectchannelphoneAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectchannelphoneAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.channelphone"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectchannelphoneAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectchannelphoneAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelPhoneDto is ChannelPhoneDto Setter
// 渠道电话对象
func (r *AlibabaalihousenewhomeprojectchannelphoneAPIRequest) SetChannelPhoneDto(_channelPhoneDto *ProjectChannelPhoneDto) error {
	r._channelPhoneDto = _channelPhoneDto
	r.Set("channel_phone_dto", _channelPhoneDto)
	return nil
}

// GetChannelPhoneDto ChannelPhoneDto Getter
func (r AlibabaalihousenewhomeprojectchannelphoneAPIRequest) GetChannelPhoneDto() *ProjectChannelPhoneDto {
	return r._channelPhoneDto
}
