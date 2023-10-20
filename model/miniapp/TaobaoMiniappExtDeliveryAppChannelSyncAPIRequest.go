package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappExtDeliveryAppChannelSyncAPIRequest ISV写入应用的渠道信息 API请求
// taobao.miniapp.ext.delivery.app.channel.sync
//
// ISV写入应用的渠道信息
type TaobaoMiniappExtDeliveryAppChannelSyncAPIRequest struct {
	model.Params
	// body
	_appChannelConfigDto *AppChannelConfigDto
}

// NewTaobaoMiniappExtDeliveryAppChannelSyncRequest 初始化TaobaoMiniappExtDeliveryAppChannelSyncAPIRequest对象
func NewTaobaoMiniappExtDeliveryAppChannelSyncRequest() *TaobaoMiniappExtDeliveryAppChannelSyncAPIRequest {
	return &TaobaoMiniappExtDeliveryAppChannelSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappExtDeliveryAppChannelSyncAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.ext.delivery.app.channel.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappExtDeliveryAppChannelSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappExtDeliveryAppChannelSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppChannelConfigDto is AppChannelConfigDto Setter
// body
func (r *TaobaoMiniappExtDeliveryAppChannelSyncAPIRequest) SetAppChannelConfigDto(_appChannelConfigDto *AppChannelConfigDto) error {
	r._appChannelConfigDto = _appChannelConfigDto
	r.Set("app_channel_config_dto", _appChannelConfigDto)
	return nil
}

// GetAppChannelConfigDto AppChannelConfigDto Getter
func (r TaobaoMiniappExtDeliveryAppChannelSyncAPIRequest) GetAppChannelConfigDto() *AppChannelConfigDto {
	return r._appChannelConfigDto
}
