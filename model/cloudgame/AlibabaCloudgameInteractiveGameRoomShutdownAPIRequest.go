package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameRoomShutdownAPIRequest 强制关闭房间 API请求
// alibaba.cloudgame.interactive.game.room.shutdown
//
// 强制关闭房间
type AlibabaCloudgameInteractiveGameRoomShutdownAPIRequest struct {
	model.Params
	// 请求入参
	_shutdownRoomRequest *ShutdownRoomRequest
}

// NewAlibabaCloudgameInteractiveGameRoomShutdownRequest 初始化AlibabaCloudgameInteractiveGameRoomShutdownAPIRequest对象
func NewAlibabaCloudgameInteractiveGameRoomShutdownRequest() *AlibabaCloudgameInteractiveGameRoomShutdownAPIRequest {
	return &AlibabaCloudgameInteractiveGameRoomShutdownAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGameRoomShutdownAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.room.shutdown"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGameRoomShutdownAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetShutdownRoomRequest is ShutdownRoomRequest Setter
// 请求入参
func (r *AlibabaCloudgameInteractiveGameRoomShutdownAPIRequest) SetShutdownRoomRequest(_shutdownRoomRequest *ShutdownRoomRequest) error {
	r._shutdownRoomRequest = _shutdownRoomRequest
	r.Set("shutdown_room_request", _shutdownRoomRequest)
	return nil
}

// GetShutdownRoomRequest ShutdownRoomRequest Getter
func (r AlibabaCloudgameInteractiveGameRoomShutdownAPIRequest) GetShutdownRoomRequest() *ShutdownRoomRequest {
	return r._shutdownRoomRequest
}
