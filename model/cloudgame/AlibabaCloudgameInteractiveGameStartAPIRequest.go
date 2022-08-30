package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameStartAPIRequest 云游戏场景互动开始游戏 API请求
// alibaba.cloudgame.interactive.game.start
//
// 开始游戏
type AlibabaCloudgameInteractiveGameStartAPIRequest struct {
	model.Params
	// 请求入参
	_startGameRequest *StartGameRequest
}

// NewAlibabaCloudgameInteractiveGameStartRequest 初始化AlibabaCloudgameInteractiveGameStartAPIRequest对象
func NewAlibabaCloudgameInteractiveGameStartRequest() *AlibabaCloudgameInteractiveGameStartAPIRequest {
	return &AlibabaCloudgameInteractiveGameStartAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGameStartAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.start"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGameStartAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStartGameRequest is StartGameRequest Setter
// 请求入参
func (r *AlibabaCloudgameInteractiveGameStartAPIRequest) SetStartGameRequest(_startGameRequest *StartGameRequest) error {
	r._startGameRequest = _startGameRequest
	r.Set("start_game_request", _startGameRequest)
	return nil
}

// GetStartGameRequest StartGameRequest Getter
func (r AlibabaCloudgameInteractiveGameStartAPIRequest) GetStartGameRequest() *StartGameRequest {
	return r._startGameRequest
}
