package cloudgame

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameGamepadGetAPIRequest 获取虚拟手柄配置 API请求
// alibaba.cloudgame.interactive.game.gamepad.get
//
// 获取虚拟手柄配置
type AlibabaCloudgameInteractiveGameGamepadGetAPIRequest struct {
	model.Params
	// 获取手柄请求入参
	_gamepadGetRequest *GamepadGetRequest
}

// NewAlibabaCloudgameInteractiveGameGamepadGetRequest 初始化AlibabaCloudgameInteractiveGameGamepadGetAPIRequest对象
func NewAlibabaCloudgameInteractiveGameGamepadGetRequest() *AlibabaCloudgameInteractiveGameGamepadGetAPIRequest {
	return &AlibabaCloudgameInteractiveGameGamepadGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCloudgameInteractiveGameGamepadGetAPIRequest) Reset() {
	r._gamepadGetRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGameGamepadGetAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.gamepad.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGameGamepadGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCloudgameInteractiveGameGamepadGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGamepadGetRequest is GamepadGetRequest Setter
// 获取手柄请求入参
func (r *AlibabaCloudgameInteractiveGameGamepadGetAPIRequest) SetGamepadGetRequest(_gamepadGetRequest *GamepadGetRequest) error {
	r._gamepadGetRequest = _gamepadGetRequest
	r.Set("gamepad_get_request", _gamepadGetRequest)
	return nil
}

// GetGamepadGetRequest GamepadGetRequest Getter
func (r AlibabaCloudgameInteractiveGameGamepadGetAPIRequest) GetGamepadGetRequest() *GamepadGetRequest {
	return r._gamepadGetRequest
}

var poolAlibabaCloudgameInteractiveGameGamepadGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCloudgameInteractiveGameGamepadGetRequest()
	},
}

// GetAlibabaCloudgameInteractiveGameGamepadGetRequest 从 sync.Pool 获取 AlibabaCloudgameInteractiveGameGamepadGetAPIRequest
func GetAlibabaCloudgameInteractiveGameGamepadGetAPIRequest() *AlibabaCloudgameInteractiveGameGamepadGetAPIRequest {
	return poolAlibabaCloudgameInteractiveGameGamepadGetAPIRequest.Get().(*AlibabaCloudgameInteractiveGameGamepadGetAPIRequest)
}

// ReleaseAlibabaCloudgameInteractiveGameGamepadGetAPIRequest 将 AlibabaCloudgameInteractiveGameGamepadGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGameGamepadGetAPIRequest(v *AlibabaCloudgameInteractiveGameGamepadGetAPIRequest) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGameGamepadGetAPIRequest.Put(v)
}
