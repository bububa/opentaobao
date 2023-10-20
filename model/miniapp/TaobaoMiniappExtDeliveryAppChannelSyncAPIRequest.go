package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappextdeliveryappchannelsyncAPIRequest ISV写入应用的渠道信息 API请求
// taobao.miniapp.ext.delivery.app.channel.sync
//
// ISV写入应用的渠道信息
type TaobaominiappextdeliveryappchannelsyncAPIRequest struct {
	model.Params
	// body
	_appChannelConfigDto *AppChannelConfigDto
}

// NewTaobaominiappextdeliveryappchannelsyncRequest 初始化TaobaominiappextdeliveryappchannelsyncAPIRequest对象
func NewTaobaominiappextdeliveryappchannelsyncRequest() *TaobaominiappextdeliveryappchannelsyncAPIRequest {
	return &TaobaominiappextdeliveryappchannelsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappextdeliveryappchannelsyncAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.ext.delivery.app.channel.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappextdeliveryappchannelsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappextdeliveryappchannelsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppChannelConfigDto is AppChannelConfigDto Setter
// body
func (r *TaobaominiappextdeliveryappchannelsyncAPIRequest) SetAppChannelConfigDto(_appChannelConfigDto *AppChannelConfigDto) error {
	r._appChannelConfigDto = _appChannelConfigDto
	r.Set("app_channel_config_dto", _appChannelConfigDto)
	return nil
}

// GetAppChannelConfigDto AppChannelConfigDto Getter
func (r TaobaominiappextdeliveryappchannelsyncAPIRequest) GetAppChannelConfigDto() *AppChannelConfigDto {
	return r._appChannelConfigDto
}
