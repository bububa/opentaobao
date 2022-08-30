package cloudgame

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGamePlayerStopAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.player.stop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGamePlayerStopAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
