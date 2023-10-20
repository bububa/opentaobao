package cloudgame

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest) Reset() {
	r._userGameStatusGetRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.player.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCloudgameInteractiveGamePlayerStatusGetRequest()
	},
}

// GetAlibabaCloudgameInteractiveGamePlayerStatusGetRequest 从 sync.Pool 获取 AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest
func GetAlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest() *AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest {
	return poolAlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest.Get().(*AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest)
}

// ReleaseAlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest 将 AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest(v *AlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGamePlayerStatusGetAPIRequest.Put(v)
}
