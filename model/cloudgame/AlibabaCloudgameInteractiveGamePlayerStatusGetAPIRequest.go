package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest 获取用户状态 API请求
// alibaba.cloudgame.interactive.game.player.status.get
//
// 获取用户状态
type AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest struct {
	model.Params
	// 请求入参
	_userGameStatusGetRequest *UserGameStatusGetRequest
}

// NewAlibabaCloudgameInteractiveGamePlayerStatusGetRequest 初始化AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest对象
func NewAlibabaCloudgameInteractiveGamePlayerStatusGetRequest() *AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest {
	return &AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.player.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUserGameStatusGetRequest is UserGameStatusGetRequest Setter
// 请求入参
func (r *AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest) SetUserGameStatusGetRequest(_userGameStatusGetRequest *UserGameStatusGetRequest) error {
	r._userGameStatusGetRequest = _userGameStatusGetRequest
	r.Set("user_game_status_get_request", _userGameStatusGetRequest)
	return nil
}

// GetUserGameStatusGetRequest UserGameStatusGetRequest Getter
func (r AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest) GetUserGameStatusGetRequest() *UserGameStatusGetRequest {
	return r._userGameStatusGetRequest
}
