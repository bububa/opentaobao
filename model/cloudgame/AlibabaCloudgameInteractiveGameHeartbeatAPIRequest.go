package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameinteractivegameheartbeatAPIRequest 游戏玩家心跳 API请求
// alibaba.cloudgame.interactive.game.heartbeat
//
// 游戏玩家心跳
type AlibabacloudgameinteractivegameheartbeatAPIRequest struct {
	model.Params
	// 请求入参
	_startGameRequest *HeartBeatRequest
}

// NewAlibabacloudgameinteractivegameheartbeatRequest 初始化AlibabacloudgameinteractivegameheartbeatAPIRequest对象
func NewAlibabacloudgameinteractivegameheartbeatRequest() *AlibabacloudgameinteractivegameheartbeatAPIRequest {
	return &AlibabacloudgameinteractivegameheartbeatAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacloudgameinteractivegameheartbeatAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.heartbeat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacloudgameinteractivegameheartbeatAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacloudgameinteractivegameheartbeatAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartGameRequest is StartGameRequest Setter
// 请求入参
func (r *AlibabacloudgameinteractivegameheartbeatAPIRequest) SetStartGameRequest(_startGameRequest *HeartBeatRequest) error {
	r._startGameRequest = _startGameRequest
	r.Set("start_game_request", _startGameRequest)
	return nil
}

// GetStartGameRequest StartGameRequest Getter
func (r AlibabacloudgameinteractivegameheartbeatAPIRequest) GetStartGameRequest() *HeartBeatRequest {
	return r._startGameRequest
}
