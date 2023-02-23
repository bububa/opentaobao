package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGamePlayerKickoutAPIRequest 踢人 API请求
// alibaba.cloudgame.interactive.game.player.kickout
//
// 踢人
type AlibabaCloudgameInteractiveGamePlayerKickoutAPIRequest struct {
	model.Params
	// 请求入参
	_kickOutUserRequest *KickOutUserRequest
}

// NewAlibabaCloudgameInteractiveGamePlayerKickoutRequest 初始化AlibabaCloudgameInteractiveGamePlayerKickoutAPIRequest对象
func NewAlibabaCloudgameInteractiveGamePlayerKickoutRequest() *AlibabaCloudgameInteractiveGamePlayerKickoutAPIRequest {
	return &AlibabaCloudgameInteractiveGamePlayerKickoutAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGamePlayerKickoutAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.player.kickout"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGamePlayerKickoutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCloudgameInteractiveGamePlayerKickoutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKickOutUserRequest is KickOutUserRequest Setter
// 请求入参
func (r *AlibabaCloudgameInteractiveGamePlayerKickoutAPIRequest) SetKickOutUserRequest(_kickOutUserRequest *KickOutUserRequest) error {
	r._kickOutUserRequest = _kickOutUserRequest
	r.Set("kick_out_user_request", _kickOutUserRequest)
	return nil
}

// GetKickOutUserRequest KickOutUserRequest Getter
func (r AlibabaCloudgameInteractiveGamePlayerKickoutAPIRequest) GetKickOutUserRequest() *KickOutUserRequest {
	return r._kickOutUserRequest
}
