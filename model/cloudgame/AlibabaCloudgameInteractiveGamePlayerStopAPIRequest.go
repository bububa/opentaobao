package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameinteractivegameplayerstopAPIRequest 用户停止游戏 API请求
// alibaba.cloudgame.interactive.game.player.stop
//
// 用户停止游戏
type AlibabacloudgameinteractivegameplayerstopAPIRequest struct {
	model.Params
	// 请求入参
	_stopUserGameRequest *StopUserGameRequest
}

// NewAlibabacloudgameinteractivegameplayerstopRequest 初始化AlibabacloudgameinteractivegameplayerstopAPIRequest对象
func NewAlibabacloudgameinteractivegameplayerstopRequest() *AlibabacloudgameinteractivegameplayerstopAPIRequest {
	return &AlibabacloudgameinteractivegameplayerstopAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacloudgameinteractivegameplayerstopAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.player.stop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacloudgameinteractivegameplayerstopAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacloudgameinteractivegameplayerstopAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStopUserGameRequest is StopUserGameRequest Setter
// 请求入参
func (r *AlibabacloudgameinteractivegameplayerstopAPIRequest) SetStopUserGameRequest(_stopUserGameRequest *StopUserGameRequest) error {
	r._stopUserGameRequest = _stopUserGameRequest
	r.Set("stop_user_game_request", _stopUserGameRequest)
	return nil
}

// GetStopUserGameRequest StopUserGameRequest Getter
func (r AlibabacloudgameinteractivegameplayerstopAPIRequest) GetStopUserGameRequest() *StopUserGameRequest {
	return r._stopUserGameRequest
}
