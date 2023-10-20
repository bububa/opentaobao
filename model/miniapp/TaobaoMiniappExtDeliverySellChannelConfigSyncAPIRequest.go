package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappextdeliverysellchannelconfigsyncAPIRequest 写入商家配置信息 API请求
// taobao.miniapp.ext.delivery.sell.channel.config.sync
//
// 写入商家配置信息
type TaobaominiappextdeliverysellchannelconfigsyncAPIRequest struct {
	model.Params
	// 入参
	_sellerChannelConfig *SellerChannelConfigDto
}

// NewTaobaominiappextdeliverysellchannelconfigsyncRequest 初始化TaobaominiappextdeliverysellchannelconfigsyncAPIRequest对象
func NewTaobaominiappextdeliverysellchannelconfigsyncRequest() *TaobaominiappextdeliverysellchannelconfigsyncAPIRequest {
	return &TaobaominiappextdeliverysellchannelconfigsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappextdeliverysellchannelconfigsyncAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.ext.delivery.sell.channel.config.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappextdeliverysellchannelconfigsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappextdeliverysellchannelconfigsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerChannelConfig is SellerChannelConfig Setter
// 入参
func (r *TaobaominiappextdeliverysellchannelconfigsyncAPIRequest) SetSellerChannelConfig(_sellerChannelConfig *SellerChannelConfigDto) error {
	r._sellerChannelConfig = _sellerChannelConfig
	r.Set("seller_channel_config", _sellerChannelConfig)
	return nil
}

// GetSellerChannelConfig SellerChannelConfig Getter
func (r TaobaominiappextdeliverysellchannelconfigsyncAPIRequest) GetSellerChannelConfig() *SellerChannelConfigDto {
	return r._sellerChannelConfig
}
