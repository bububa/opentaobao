package cloudgame

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameHeartbeatAPIRequest 游戏玩家心跳 API请求
// alibaba.cloudgame.interactive.game.heartbeat
//
// 游戏玩家心跳
type AlibabaCloudgameInteractiveGameHeartbeatAPIRequest struct {
	model.Params
	// 请求入参
	_startGameRequest *HeartBeatRequest
}

// NewAlibabaCloudgameInteractiveGameHeartbeatRequest 初始化AlibabaCloudgameInteractiveGameHeartbeatAPIRequest对象
func NewAlibabaCloudgameInteractiveGameHeartbeatRequest() *AlibabaCloudgameInteractiveGameHeartbeatAPIRequest {
	return &AlibabaCloudgameInteractiveGameHeartbeatAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCloudgameInteractiveGameHeartbeatAPIRequest) Reset() {
	r._startGameRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGameHeartbeatAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.heartbeat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGameHeartbeatAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCloudgameInteractiveGameHeartbeatAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartGameRequest is StartGameRequest Setter
// 请求入参
func (r *AlibabaCloudgameInteractiveGameHeartbeatAPIRequest) SetStartGameRequest(_startGameRequest *HeartBeatRequest) error {
	r._startGameRequest = _startGameRequest
	r.Set("start_game_request", _startGameRequest)
	return nil
}

// GetStartGameRequest StartGameRequest Getter
func (r AlibabaCloudgameInteractiveGameHeartbeatAPIRequest) GetStartGameRequest() *HeartBeatRequest {
	return r._startGameRequest
}

var poolAlibabaCloudgameInteractiveGameHeartbeatAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCloudgameInteractiveGameHeartbeatRequest()
	},
}

// GetAlibabaCloudgameInteractiveGameHeartbeatRequest 从 sync.Pool 获取 AlibabaCloudgameInteractiveGameHeartbeatAPIRequest
func GetAlibabaCloudgameInteractiveGameHeartbeatAPIRequest() *AlibabaCloudgameInteractiveGameHeartbeatAPIRequest {
	return poolAlibabaCloudgameInteractiveGameHeartbeatAPIRequest.Get().(*AlibabaCloudgameInteractiveGameHeartbeatAPIRequest)
}

// ReleaseAlibabaCloudgameInteractiveGameHeartbeatAPIRequest 将 AlibabaCloudgameInteractiveGameHeartbeatAPIRequest 放入 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGameHeartbeatAPIRequest(v *AlibabaCloudgameInteractiveGameHeartbeatAPIRequest) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGameHeartbeatAPIRequest.Put(v)
}
