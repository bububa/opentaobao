package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCsrGameDataSyncCheckAPIRequest 公益互动 外部游戏数据同步-校验 API请求
// alibaba.csr.game.data.sync.check
//
// 公益互动 外部游戏数据同步-校验
type AlibabaCsrGameDataSyncCheckAPIRequest struct {
	model.Params
	// 请求
	_snakeDataCheckRequest *SnakeDataCheckRequest
}

// NewAlibabaCsrGameDataSyncCheckRequest 初始化AlibabaCsrGameDataSyncCheckAPIRequest对象
func NewAlibabaCsrGameDataSyncCheckRequest() *AlibabaCsrGameDataSyncCheckAPIRequest {
	return &AlibabaCsrGameDataSyncCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCsrGameDataSyncCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.csr.game.data.sync.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCsrGameDataSyncCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCsrGameDataSyncCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSnakeDataCheckRequest is SnakeDataCheckRequest Setter
// 请求
func (r *AlibabaCsrGameDataSyncCheckAPIRequest) SetSnakeDataCheckRequest(_snakeDataCheckRequest *SnakeDataCheckRequest) error {
	r._snakeDataCheckRequest = _snakeDataCheckRequest
	r.Set("snake_data_check_request", _snakeDataCheckRequest)
	return nil
}

// GetSnakeDataCheckRequest SnakeDataCheckRequest Getter
func (r AlibabaCsrGameDataSyncCheckAPIRequest) GetSnakeDataCheckRequest() *SnakeDataCheckRequest {
	return r._snakeDataCheckRequest
}
