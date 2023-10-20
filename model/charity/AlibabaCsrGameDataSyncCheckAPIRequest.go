package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacsrgamedatasynccheckAPIRequest 公益互动 外部游戏数据同步-校验 API请求
// alibaba.csr.game.data.sync.check
//
// 公益互动 外部游戏数据同步-校验
type AlibabacsrgamedatasynccheckAPIRequest struct {
	model.Params
	// 请求
	_snakeDataCheckRequest *SnakeDataCheckRequest
}

// NewAlibabacsrgamedatasynccheckRequest 初始化AlibabacsrgamedatasynccheckAPIRequest对象
func NewAlibabacsrgamedatasynccheckRequest() *AlibabacsrgamedatasynccheckAPIRequest {
	return &AlibabacsrgamedatasynccheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacsrgamedatasynccheckAPIRequest) GetApiMethodName() string {
	return "alibaba.csr.game.data.sync.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacsrgamedatasynccheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacsrgamedatasynccheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSnakeDataCheckRequest is SnakeDataCheckRequest Setter
// 请求
func (r *AlibabacsrgamedatasynccheckAPIRequest) SetSnakeDataCheckRequest(_snakeDataCheckRequest *SnakeDataCheckRequest) error {
	r._snakeDataCheckRequest = _snakeDataCheckRequest
	r.Set("snake_data_check_request", _snakeDataCheckRequest)
	return nil
}

// GetSnakeDataCheckRequest SnakeDataCheckRequest Getter
func (r AlibabacsrgamedatasynccheckAPIRequest) GetSnakeDataCheckRequest() *SnakeDataCheckRequest {
	return r._snakeDataCheckRequest
}
