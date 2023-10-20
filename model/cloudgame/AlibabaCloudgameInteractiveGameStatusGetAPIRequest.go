package cloudgame

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameStatusGetAPIRequest 获取游戏状态 API请求
// alibaba.cloudgame.interactive.game.status.get
//
// 获取游戏状态
type AlibabaCloudgameInteractiveGameStatusGetAPIRequest struct {
	model.Params
	// 请求入参
	_gameStatusGetRequest *GameStatusGetRequest
}

// NewAlibabaCloudgameInteractiveGameStatusGetRequest 初始化AlibabaCloudgameInteractiveGameStatusGetAPIRequest对象
func NewAlibabaCloudgameInteractiveGameStatusGetRequest() *AlibabaCloudgameInteractiveGameStatusGetAPIRequest {
	return &AlibabaCloudgameInteractiveGameStatusGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCloudgameInteractiveGameStatusGetAPIRequest) Reset() {
	r._gameStatusGetRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGameStatusGetAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGameStatusGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCloudgameInteractiveGameStatusGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGameStatusGetRequest is GameStatusGetRequest Setter
// 请求入参
func (r *AlibabaCloudgameInteractiveGameStatusGetAPIRequest) SetGameStatusGetRequest(_gameStatusGetRequest *GameStatusGetRequest) error {
	r._gameStatusGetRequest = _gameStatusGetRequest
	r.Set("game_status_get_request", _gameStatusGetRequest)
	return nil
}

// GetGameStatusGetRequest GameStatusGetRequest Getter
func (r AlibabaCloudgameInteractiveGameStatusGetAPIRequest) GetGameStatusGetRequest() *GameStatusGetRequest {
	return r._gameStatusGetRequest
}

var poolAlibabaCloudgameInteractiveGameStatusGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCloudgameInteractiveGameStatusGetRequest()
	},
}

// GetAlibabaCloudgameInteractiveGameStatusGetRequest 从 sync.Pool 获取 AlibabaCloudgameInteractiveGameStatusGetAPIRequest
func GetAlibabaCloudgameInteractiveGameStatusGetAPIRequest() *AlibabaCloudgameInteractiveGameStatusGetAPIRequest {
	return poolAlibabaCloudgameInteractiveGameStatusGetAPIRequest.Get().(*AlibabaCloudgameInteractiveGameStatusGetAPIRequest)
}

// ReleaseAlibabaCloudgameInteractiveGameStatusGetAPIRequest 将 AlibabaCloudgameInteractiveGameStatusGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGameStatusGetAPIRequest(v *AlibabaCloudgameInteractiveGameStatusGetAPIRequest) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGameStatusGetAPIRequest.Put(v)
}
