package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameinteractivegamegamepadgetAPIRequest 获取虚拟手柄配置 API请求
// alibaba.cloudgame.interactive.game.gamepad.get
//
// 获取虚拟手柄配置
type AlibabacloudgameinteractivegamegamepadgetAPIRequest struct {
	model.Params
	// 获取手柄请求入参
	_gamepadGetRequest *GamepadGetRequest
}

// NewAlibabacloudgameinteractivegamegamepadgetRequest 初始化AlibabacloudgameinteractivegamegamepadgetAPIRequest对象
func NewAlibabacloudgameinteractivegamegamepadgetRequest() *AlibabacloudgameinteractivegamegamepadgetAPIRequest {
	return &AlibabacloudgameinteractivegamegamepadgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacloudgameinteractivegamegamepadgetAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.gamepad.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacloudgameinteractivegamegamepadgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacloudgameinteractivegamegamepadgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGamepadGetRequest is GamepadGetRequest Setter
// 获取手柄请求入参
func (r *AlibabacloudgameinteractivegamegamepadgetAPIRequest) SetGamepadGetRequest(_gamepadGetRequest *GamepadGetRequest) error {
	r._gamepadGetRequest = _gamepadGetRequest
	r.Set("gamepad_get_request", _gamepadGetRequest)
	return nil
}

// GetGamepadGetRequest GamepadGetRequest Getter
func (r AlibabacloudgameinteractivegamegamepadgetAPIRequest) GetGamepadGetRequest() *GamepadGetRequest {
	return r._gamepadGetRequest
}
