package cloudgame

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameStartAPIRequest 云游戏场景互动开始游戏 API请求
// alibaba.cloudgame.interactive.game.start
//
// 开始游戏
type AlibabaCloudgameInteractiveGameStartAPIRequest struct {
	model.Params
	// 请求入参
	_startGameRequest *StartGameRequest
}

// NewAlibabaCloudgameInteractiveGameStartRequest 初始化AlibabaCloudgameInteractiveGameStartAPIRequest对象
func NewAlibabaCloudgameInteractiveGameStartRequest() *AlibabaCloudgameInteractiveGameStartAPIRequest {
	return &AlibabaCloudgameInteractiveGameStartAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCloudgameInteractiveGameStartAPIRequest) Reset() {
	r._startGameRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameInteractiveGameStartAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.interactive.game.start"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameInteractiveGameStartAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCloudgameInteractiveGameStartAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartGameRequest is StartGameRequest Setter
// 请求入参
func (r *AlibabaCloudgameInteractiveGameStartAPIRequest) SetStartGameRequest(_startGameRequest *StartGameRequest) error {
	r._startGameRequest = _startGameRequest
	r.Set("start_game_request", _startGameRequest)
	return nil
}

// GetStartGameRequest StartGameRequest Getter
func (r AlibabaCloudgameInteractiveGameStartAPIRequest) GetStartGameRequest() *StartGameRequest {
	return r._startGameRequest
}

var poolAlibabaCloudgameInteractiveGameStartAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCloudgameInteractiveGameStartRequest()
	},
}

// GetAlibabaCloudgameInteractiveGameStartRequest 从 sync.Pool 获取 AlibabaCloudgameInteractiveGameStartAPIRequest
func GetAlibabaCloudgameInteractiveGameStartAPIRequest() *AlibabaCloudgameInteractiveGameStartAPIRequest {
	return poolAlibabaCloudgameInteractiveGameStartAPIRequest.Get().(*AlibabaCloudgameInteractiveGameStartAPIRequest)
}

// ReleaseAlibabaCloudgameInteractiveGameStartAPIRequest 将 AlibabaCloudgameInteractiveGameStartAPIRequest 放入 sync.Pool
func ReleaseAlibabaCloudgameInteractiveGameStartAPIRequest(v *AlibabaCloudgameInteractiveGameStartAPIRequest) {
	v.Reset()
	poolAlibabaCloudgameInteractiveGameStartAPIRequest.Put(v)
}
