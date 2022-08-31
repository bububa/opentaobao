package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectChannelphoneAPIRequest 新房渠道电话数据同步 API请求
// alibaba.alihouse.newhome.project.channelphone
//
// 新房渠道电话数据同步
type AlibabaAlihouseNewhomeProjectChannelphoneAPIRequest struct {
	model.Params
	// 渠道电话对象
	_channelPhoneDto *ProjectChannelPhoneDto
}

// NewAlibabaAlihouseNewhomeProjectChannelphoneRequest 初始化AlibabaAlihouseNewhomeProjectChannelphoneAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectChannelphoneRequest() *AlibabaAlihouseNewhomeProjectChannelphoneAPIRequest {
	return &AlibabaAlihouseNewhomeProjectChannelphoneAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectChannelphoneAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.channelphone"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectChannelphoneAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetChannelPhoneDto is ChannelPhoneDto Setter
// 渠道电话对象
func (r *AlibabaAlihouseNewhomeProjectChannelphoneAPIRequest) SetChannelPhoneDto(_channelPhoneDto *ProjectChannelPhoneDto) error {
	r._channelPhoneDto = _channelPhoneDto
	r.Set("channel_phone_dto", _channelPhoneDto)
	return nil
}

// GetChannelPhoneDto ChannelPhoneDto Getter
func (r AlibabaAlihouseNewhomeProjectChannelphoneAPIRequest) GetChannelPhoneDto() *ProjectChannelPhoneDto {
	return r._channelPhoneDto
}
