package cloudgame

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGameHeartbeatAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.heartbeat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGameHeartbeatAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
