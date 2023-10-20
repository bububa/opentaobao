package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacsrgamedatasyncAPIRequest 公益互动 外部游戏数据同步 API请求
// alibaba.csr.game.data.sync
//
// 公益互动 外部游戏数据同步
type AlibabacsrgamedatasyncAPIRequest struct {
	model.Params
	// 请求
	_snakeDataSyncRequest *SnakeDataSyncRequest
}

// NewAlibabacsrgamedatasyncRequest 初始化AlibabacsrgamedatasyncAPIRequest对象
func NewAlibabacsrgamedatasyncRequest() *AlibabacsrgamedatasyncAPIRequest {
	return &AlibabacsrgamedatasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacsrgamedatasyncAPIRequest) GetApiMethodName() string {
	return "alibaba.csr.game.data.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacsrgamedatasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacsrgamedatasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSnakeDataSyncRequest is SnakeDataSyncRequest Setter
// 请求
func (r *AlibabacsrgamedatasyncAPIRequest) SetSnakeDataSyncRequest(_snakeDataSyncRequest *SnakeDataSyncRequest) error {
	r._snakeDataSyncRequest = _snakeDataSyncRequest
	r.Set("snake_data_sync_request", _snakeDataSyncRequest)
	return nil
}

// GetSnakeDataSyncRequest SnakeDataSyncRequest Getter
func (r AlibabacsrgamedatasyncAPIRequest) GetSnakeDataSyncRequest() *SnakeDataSyncRequest {
	return r._snakeDataSyncRequest
}
