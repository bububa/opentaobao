package tmallhk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallHkClearanceOrderGetAPIRequest 天猫国际订单清关信息 API请求
// tmall.hk.clearance.order.get
//
// 天猫国际订单清关信息
type TmallHkClearanceOrderGetAPIRequest struct {
	model.Params
	// 入参封装类型
	_clearanceOrderRequest *ClearanceOrderRequest
}

// NewTmallHkClearanceOrderGetRequest 初始化TmallHkClearanceOrderGetAPIRequest对象
func NewTmallHkClearanceOrderGetRequest() *TmallHkClearanceOrderGetAPIRequest {
	return &TmallHkClearanceOrderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallHkClearanceOrderGetAPIRequest) Reset() {
	r._clearanceOrderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallHkClearanceOrderGetAPIRequest) GetApiMethodName() string {
	return "tmall.hk.clearance.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallHkClearanceOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallHkClearanceOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClearanceOrderRequest is ClearanceOrderRequest Setter
// 入参封装类型
func (r *TmallHkClearanceOrderGetAPIRequest) SetClearanceOrderRequest(_clearanceOrderRequest *ClearanceOrderRequest) error {
	r._clearanceOrderRequest = _clearanceOrderRequest
	r.Set("clearance_order_request", _clearanceOrderRequest)
	return nil
}

// GetClearanceOrderRequest ClearanceOrderRequest Getter
func (r TmallHkClearanceOrderGetAPIRequest) GetClearanceOrderRequest() *ClearanceOrderRequest {
	return r._clearanceOrderRequest
}

var poolTmallHkClearanceOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallHkClearanceOrderGetRequest()
	},
}

// GetTmallHkClearanceOrderGetRequest 从 sync.Pool 获取 TmallHkClearanceOrderGetAPIRequest
func GetTmallHkClearanceOrderGetAPIRequest() *TmallHkClearanceOrderGetAPIRequest {
	return poolTmallHkClearanceOrderGetAPIRequest.Get().(*TmallHkClearanceOrderGetAPIRequest)
}

// ReleaseTmallHkClearanceOrderGetAPIRequest 将 TmallHkClearanceOrderGetAPIRequest 放入 sync.Pool
func ReleaseTmallHkClearanceOrderGetAPIRequest(v *TmallHkClearanceOrderGetAPIRequest) {
	v.Reset()
	poolTmallHkClearanceOrderGetAPIRequest.Put(v)
}
