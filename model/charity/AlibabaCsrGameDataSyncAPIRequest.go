package charity

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCsrGameDataSyncAPIRequest 公益互动 外部游戏数据同步 API请求
// alibaba.csr.game.data.sync
//
// 公益互动 外部游戏数据同步
type AlibabaCsrGameDataSyncAPIRequest struct {
	model.Params
	// 请求
	_snakeDataSyncRequest *SnakeDataSyncRequest
}

// NewAlibabaCsrGameDataSyncRequest 初始化AlibabaCsrGameDataSyncAPIRequest对象
func NewAlibabaCsrGameDataSyncRequest() *AlibabaCsrGameDataSyncAPIRequest {
	return &AlibabaCsrGameDataSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCsrGameDataSyncAPIRequest) Reset() {
	r._snakeDataSyncRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCsrGameDataSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.csr.game.data.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCsrGameDataSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCsrGameDataSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSnakeDataSyncRequest is SnakeDataSyncRequest Setter
// 请求
func (r *AlibabaCsrGameDataSyncAPIRequest) SetSnakeDataSyncRequest(_snakeDataSyncRequest *SnakeDataSyncRequest) error {
	r._snakeDataSyncRequest = _snakeDataSyncRequest
	r.Set("snake_data_sync_request", _snakeDataSyncRequest)
	return nil
}

// GetSnakeDataSyncRequest SnakeDataSyncRequest Getter
func (r AlibabaCsrGameDataSyncAPIRequest) GetSnakeDataSyncRequest() *SnakeDataSyncRequest {
	return r._snakeDataSyncRequest
}

var poolAlibabaCsrGameDataSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCsrGameDataSyncRequest()
	},
}

// GetAlibabaCsrGameDataSyncRequest 从 sync.Pool 获取 AlibabaCsrGameDataSyncAPIRequest
func GetAlibabaCsrGameDataSyncAPIRequest() *AlibabaCsrGameDataSyncAPIRequest {
	return poolAlibabaCsrGameDataSyncAPIRequest.Get().(*AlibabaCsrGameDataSyncAPIRequest)
}

// ReleaseAlibabaCsrGameDataSyncAPIRequest 将 AlibabaCsrGameDataSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaCsrGameDataSyncAPIRequest(v *AlibabaCsrGameDataSyncAPIRequest) {
	v.Reset()
	poolAlibabaCsrGameDataSyncAPIRequest.Put(v)
}
