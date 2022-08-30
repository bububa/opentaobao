package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameRoomCreateAPIRequest 建游戏房间 API请求
// alibaba.cloudgame.interactive.game.room.create
//
// 建游戏房间
type AlibabaCloudgameInteractiveGameRoomCreateAPIRequest struct {
	model.Params
	// 请求
	_createRoomRequest *CreateRoomRequest
}

// NewAlibabaCloudgameInteractiveGameRoomCreateRequest 初始化AlibabaCloudgameInteractiveGameRoomCreateAPIRequest对象
func NewAlibabaCloudgameInteractiveGameRoomCreateRequest() *AlibabaCloudgameInteractiveGameRoomCreateAPIRequest {
	return &AlibabaCloudgameInteractiveGameRoomCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGameRoomCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.room.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGameRoomCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCreateRoomRequest is CreateRoomRequest Setter
// 请求
func (r *AlibabaCloudgameInteractiveGameRoomCreateAPIRequest) SetCreateRoomRequest(_createRoomRequest *CreateRoomRequest) error {
	r._createRoomRequest = _createRoomRequest
	r.Set("create_room_request", _createRoomRequest)
	return nil
}

// GetCreateRoomRequest CreateRoomRequest Getter
func (r AlibabaCloudgameInteractiveGameRoomCreateAPIRequest) GetCreateRoomRequest() *CreateRoomRequest {
	return r._createRoomRequest
}
