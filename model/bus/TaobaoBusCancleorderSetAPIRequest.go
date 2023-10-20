package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusCancleorderSetAPIRequest 取消订单 API请求
// taobao.bus.cancleorder.set
//
// 取消订单
type TaobaoBusCancleorderSetAPIRequest struct {
	model.Params
	// 阿里订单号
	_aliOrderId string
}

// NewTaobaoBusCancleorderSetRequest 初始化TaobaoBusCancleorderSetAPIRequest对象
func NewTaobaoBusCancleorderSetRequest() *TaobaoBusCancleorderSetAPIRequest {
	return &TaobaoBusCancleorderSetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusCancleorderSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.cancleorder.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusCancleorderSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusCancleorderSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAliOrderId is AliOrderId Setter
// 阿里订单号
func (r *TaobaoBusCancleorderSetAPIRequest) SetAliOrderId(_aliOrderId string) error {
	r._aliOrderId = _aliOrderId
	r.Set("ali_order_id", _aliOrderId)
	return nil
}

// GetAliOrderId AliOrderId Getter
func (r TaobaoBusCancleorderSetAPIRequest) GetAliOrderId() string {
	return r._aliOrderId
}
