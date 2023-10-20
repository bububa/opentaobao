package cloudgame

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGamePlayerStopAPIRequest 用户停止游戏 API请求
// alibaba.cloudgame.interactive.game.player.stop
//
// 用户停止游戏
type AlibabaCloudgameInteractiveGamePlayerStopAPIRequest struct {
	model.Params
	// 请求入参
	_stopUserGameRequest *StopUserGameRequest
}

// NewAlibabaCloudgameInteractiveGamePlayerStopRequest 初始化AlibabaCloudgameInteractiveGamePlayerStopAPIRequest对象
func NewAlibabaCloudgameInteractiveGamePlayerStopRequest() *AlibabaCloudgameInteractiveGamePlayerStopAPIRequest {
	return &AlibabaCloudgameInteractiveGamePlayerStopAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCloudgameInteractiveGamePlayerStopAPIRequest) Reset() {
	r._stopUserGameRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGamePlayerStopAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.player.stop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGamePlayerStopAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCloudgameInteractiveGamePlayerStopAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStopUserGameRequest is StopUserGameRequest Setter
// 请求入参
func (r *AlibabaCloudgameInteractiveGamePlayerStopAPIRequest) SetStopUserGameRequest(_stopUserGameRequest *StopUserGameRequest) error {
	r._stopUserGameRequest = _stopUserGameRequest
	r.Set("stop_user_game_request", _stopUserGameRequest)
	return nil
}

// GetStopUserGameRequest StopUserGameRequest Getter
func (r AlibabaCloudgameInteractiveGamePlayerStopAPIRequest) GetStopUserGameRequest() *StopUserGameRequest {
	return r._stopUserGameRequest
}

var poolAlibabaCloudgameInteractiveGamePlayerStopAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCloudgameInteractiveGamePlayerStopRequest()
	},
}

// GetAlibabaCloudgameInteractiveGamePlayerStopRequest 从 sync.Pool 获取 AlibabaCloudgameInteractiveGamePlayerStopAPIRequest
func GetAlibabaCloudgameInteractiveGamePlayerStopAPIRequest() *AlibabaCloudgameInteractiveGamePlayerStopAPIRequest {
	return poolAlibabaCloudgameInteractiveGamePlayerStopAPIRequest.Get().(*AlibabaCloudgameInteractiveGamePlayerStopAPIRequest)
}

// ReleaseAlibabaCloudgameInteractiveGamePlayerStopAPIRequest 将 AlibabaCloudgameInteractiveGamePlayerStopAPIRequest 放入 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGamePlayerStopAPIRequest(v *AlibabaCloudgameInteractiveGamePlayerStopAPIRequest) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGamePlayerStopAPIRequest.Put(v)
}
