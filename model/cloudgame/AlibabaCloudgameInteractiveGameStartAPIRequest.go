package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameinteractivegamestartAPIRequest 云游戏场景互动开始游戏 API请求
// alibaba.cloudgame.interactive.game.start
//
// 开始游戏
type AlibabacloudgameinteractivegamestartAPIRequest struct {
	model.Params
	// 请求入参
	_startGameRequest *StartGameRequest
}

// NewAlibabacloudgameinteractivegamestartRequest 初始化AlibabacloudgameinteractivegamestartAPIRequest对象
func NewAlibabacloudgameinteractivegamestartRequest() *AlibabacloudgameinteractivegamestartAPIRequest {
	return &AlibabacloudgameinteractivegamestartAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacloudgameinteractivegamestartAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.start"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacloudgameinteractivegamestartAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacloudgameinteractivegamestartAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartGameRequest is StartGameRequest Setter
// 请求入参
func (r *AlibabacloudgameinteractivegamestartAPIRequest) SetStartGameRequest(_startGameRequest *StartGameRequest) error {
	r._startGameRequest = _startGameRequest
	r.Set("start_game_request", _startGameRequest)
	return nil
}

// GetStartGameRequest StartGameRequest Getter
func (r AlibabacloudgameinteractivegamestartAPIRequest) GetStartGameRequest() *StartGameRequest {
	return r._startGameRequest
}
