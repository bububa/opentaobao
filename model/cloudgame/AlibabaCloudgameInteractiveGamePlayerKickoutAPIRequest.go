package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameinteractivegameplayerkickoutAPIRequest 踢人 API请求
// alibaba.cloudgame.interactive.game.player.kickout
//
// 踢人
type AlibabacloudgameinteractivegameplayerkickoutAPIRequest struct {
	model.Params
	// 请求入参
	_kickOutUserRequest *KickOutUserRequest
}

// NewAlibabacloudgameinteractivegameplayerkickoutRequest 初始化AlibabacloudgameinteractivegameplayerkickoutAPIRequest对象
func NewAlibabacloudgameinteractivegameplayerkickoutRequest() *AlibabacloudgameinteractivegameplayerkickoutAPIRequest {
	return &AlibabacloudgameinteractivegameplayerkickoutAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacloudgameinteractivegameplayerkickoutAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.player.kickout"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacloudgameinteractivegameplayerkickoutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacloudgameinteractivegameplayerkickoutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKickOutUserRequest is KickOutUserRequest Setter
// 请求入参
func (r *AlibabacloudgameinteractivegameplayerkickoutAPIRequest) SetKickOutUserRequest(_kickOutUserRequest *KickOutUserRequest) error {
	r._kickOutUserRequest = _kickOutUserRequest
	r.Set("kick_out_user_request", _kickOutUserRequest)
	return nil
}

// GetKickOutUserRequest KickOutUserRequest Getter
func (r AlibabacloudgameinteractivegameplayerkickoutAPIRequest) GetKickOutUserRequest() *KickOutUserRequest {
	return r._kickOutUserRequest
}
