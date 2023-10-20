package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameinteractivegameroomcreateAPIRequest 建游戏房间 API请求
// alibaba.cloudgame.interactive.game.room.create
//
// 建游戏房间
type AlibabacloudgameinteractivegameroomcreateAPIRequest struct {
	model.Params
	// 请求
	_createRoomRequest *CreateRoomRequest
}

// NewAlibabacloudgameinteractivegameroomcreateRequest 初始化AlibabacloudgameinteractivegameroomcreateAPIRequest对象
func NewAlibabacloudgameinteractivegameroomcreateRequest() *AlibabacloudgameinteractivegameroomcreateAPIRequest {
	return &AlibabacloudgameinteractivegameroomcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacloudgameinteractivegameroomcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.room.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacloudgameinteractivegameroomcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacloudgameinteractivegameroomcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateRoomRequest is CreateRoomRequest Setter
// 请求
func (r *AlibabacloudgameinteractivegameroomcreateAPIRequest) SetCreateRoomRequest(_createRoomRequest *CreateRoomRequest) error {
	r._createRoomRequest = _createRoomRequest
	r.Set("create_room_request", _createRoomRequest)
	return nil
}

// GetCreateRoomRequest CreateRoomRequest Getter
func (r AlibabacloudgameinteractivegameroomcreateAPIRequest) GetCreateRoomRequest() *CreateRoomRequest {
	return r._createRoomRequest
}
