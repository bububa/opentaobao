package cloudgame

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCloudgameInteractiveGameRoomCreateAPIRequest) Reset() {
	r._createRoomRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGameRoomCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.room.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGameRoomCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCloudgameInteractiveGameRoomCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaCloudgameInteractiveGameRoomCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCloudgameInteractiveGameRoomCreateRequest()
	},
}

// GetAlibabaCloudgameInteractiveGameRoomCreateRequest 从 sync.Pool 获取 AlibabaCloudgameInteractiveGameRoomCreateAPIRequest
func GetAlibabaCloudgameInteractiveGameRoomCreateAPIRequest() *AlibabaCloudgameInteractiveGameRoomCreateAPIRequest {
	return poolAlibabaCloudgameInteractiveGameRoomCreateAPIRequest.Get().(*AlibabaCloudgameInteractiveGameRoomCreateAPIRequest)
}

// ReleaseAlibabaCloudgameInteractiveGameRoomCreateAPIRequest 将 AlibabaCloudgameInteractiveGameRoomCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGameRoomCreateAPIRequest(v *AlibabaCloudgameInteractiveGameRoomCreateAPIRequest) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGameRoomCreateAPIRequest.Put(v)
}
