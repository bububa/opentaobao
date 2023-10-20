package baodian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaodianserversdkconfiggetAPIRequest 获取宝点SDK的配置项（已迁移） API请求
// taobao.baodian.server.sdk.config.get
//
// 获取SDK各种配置项（已迁移）
type TaobaobaodianserversdkconfiggetAPIRequest struct {
	model.Params
	// appKey
	_appkey string
	// 渠道
	_channel string
	// sdk版本号
	_sdkVer string
	// 与后端约定
	_type int64
}

// NewTaobaobaodianserversdkconfiggetRequest 初始化TaobaobaodianserversdkconfiggetAPIRequest对象
func NewTaobaobaodianserversdkconfiggetRequest() *TaobaobaodianserversdkconfiggetAPIRequest {
	return &TaobaobaodianserversdkconfiggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaodianserversdkconfiggetAPIRequest) GetApiMethodName() string {
	return "taobao.baodian.server.sdk.config.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaodianserversdkconfiggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaodianserversdkconfiggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppkey is Appkey Setter
// appKey
func (r *TaobaobaodianserversdkconfiggetAPIRequest) SetAppkey(_appkey string) error {
	r._appkey = _appkey
	r.Set("appkey", _appkey)
	return nil
}

// GetAppkey Appkey Getter
func (r TaobaobaodianserversdkconfiggetAPIRequest) GetAppkey() string {
	return r._appkey
}

// SetChannel is Channel Setter
// 渠道
func (r *TaobaobaodianserversdkconfiggetAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r TaobaobaodianserversdkconfiggetAPIRequest) GetChannel() string {
	return r._channel
}

// SetSdkVer is SdkVer Setter
// sdk版本号
func (r *TaobaobaodianserversdkconfiggetAPIRequest) SetSdkVer(_sdkVer string) error {
	r._sdkVer = _sdkVer
	r.Set("sdk_ver", _sdkVer)
	return nil
}

// GetSdkVer SdkVer Getter
func (r TaobaobaodianserversdkconfiggetAPIRequest) GetSdkVer() string {
	return r._sdkVer
}

// SetType is Type Setter
// 与后端约定
func (r *TaobaobaodianserversdkconfiggetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaobaodianserversdkconfiggetAPIRequest) GetType() int64 {
	return r._type
}
