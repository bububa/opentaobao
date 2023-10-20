package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameinteractivegamestatusgetAPIRequest 获取游戏状态 API请求
// alibaba.cloudgame.interactive.game.status.get
//
// 获取游戏状态
type AlibabacloudgameinteractivegamestatusgetAPIRequest struct {
	model.Params
	// 请求入参
	_gameStatusGetRequest *GameStatusGetRequest
}

// NewAlibabacloudgameinteractivegamestatusgetRequest 初始化AlibabacloudgameinteractivegamestatusgetAPIRequest对象
func NewAlibabacloudgameinteractivegamestatusgetRequest() *AlibabacloudgameinteractivegamestatusgetAPIRequest {
	return &AlibabacloudgameinteractivegamestatusgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacloudgameinteractivegamestatusgetAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacloudgameinteractivegamestatusgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacloudgameinteractivegamestatusgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGameStatusGetRequest is GameStatusGetRequest Setter
// 请求入参
func (r *AlibabacloudgameinteractivegamestatusgetAPIRequest) SetGameStatusGetRequest(_gameStatusGetRequest *GameStatusGetRequest) error {
	r._gameStatusGetRequest = _gameStatusGetRequest
	r.Set("game_status_get_request", _gameStatusGetRequest)
	return nil
}

// GetGameStatusGetRequest GameStatusGetRequest Getter
func (r AlibabacloudgameinteractivegamestatusgetAPIRequest) GetGameStatusGetRequest() *GameStatusGetRequest {
	return r._gameStatusGetRequest
}
