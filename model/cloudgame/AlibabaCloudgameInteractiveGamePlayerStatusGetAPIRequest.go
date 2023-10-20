package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameinteractivegameplayerstatusgetAPIRequest 获取用户状态 API请求
// alibaba.cloudgame.interactive.game.player.status.get
//
// 获取用户状态
type AlibabacloudgameinteractivegameplayerstatusgetAPIRequest struct {
	model.Params
	// 请求入参
	_userGameStatusGetRequest *UserGameStatusGetRequest
}

// NewAlibabacloudgameinteractivegameplayerstatusgetRequest 初始化AlibabacloudgameinteractivegameplayerstatusgetAPIRequest对象
func NewAlibabacloudgameinteractivegameplayerstatusgetRequest() *AlibabacloudgameinteractivegameplayerstatusgetAPIRequest {
	return &AlibabacloudgameinteractivegameplayerstatusgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacloudgameinteractivegameplayerstatusgetAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.player.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacloudgameinteractivegameplayerstatusgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacloudgameinteractivegameplayerstatusgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserGameStatusGetRequest is UserGameStatusGetRequest Setter
// 请求入参
func (r *AlibabacloudgameinteractivegameplayerstatusgetAPIRequest) SetUserGameStatusGetRequest(_userGameStatusGetRequest *UserGameStatusGetRequest) error {
	r._userGameStatusGetRequest = _userGameStatusGetRequest
	r.Set("user_game_status_get_request", _userGameStatusGetRequest)
	return nil
}

// GetUserGameStatusGetRequest UserGameStatusGetRequest Getter
func (r AlibabacloudgameinteractivegameplayerstatusgetAPIRequest) GetUserGameStatusGetRequest() *UserGameStatusGetRequest {
	return r._userGameStatusGetRequest
}
