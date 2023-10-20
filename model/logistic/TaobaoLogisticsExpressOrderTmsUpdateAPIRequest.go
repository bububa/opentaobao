package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressOrderTmsUpdateAPIRequest 服务商修改上门取退时间接口 API请求
// taobao.logistics.express.order.tms.update
//
// 服务商修改上门取退时间接口
type TaobaoLogisticsExpressOrderTmsUpdateAPIRequest struct {
	model.Params
	// tms更新参数
	_tmsToMscUpdateOrderRequest *Tms2MscUpdateOrderRequest
}

// NewTaobaoLogisticsExpressOrderTmsUpdateRequest 初始化TaobaoLogisticsExpressOrderTmsUpdateAPIRequest对象
func NewTaobaoLogisticsExpressOrderTmsUpdateRequest() *TaobaoLogisticsExpressOrderTmsUpdateAPIRequest {
	return &TaobaoLogisticsExpressOrderTmsUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsExpressOrderTmsUpdateAPIRequest) Reset() {
	r._tmsToMscUpdateOrderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressOrderTmsUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.order.tms.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressOrderTmsUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressOrderTmsUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsToMscUpdateOrderRequest is TmsToMscUpdateOrderRequest Setter
// tms更新参数
func (r *TaobaoLogisticsExpressOrderTmsUpdateAPIRequest) SetTmsToMscUpdateOrderRequest(_tmsToMscUpdateOrderRequest *Tms2MscUpdateOrderRequest) error {
	r._tmsToMscUpdateOrderRequest = _tmsToMscUpdateOrderRequest
	r.Set("tms_to_msc_update_order_request", _tmsToMscUpdateOrderRequest)
	return nil
}

// GetTmsToMscUpdateOrderRequest TmsToMscUpdateOrderRequest Getter
func (r TaobaoLogisticsExpressOrderTmsUpdateAPIRequest) GetTmsToMscUpdateOrderRequest() *Tms2MscUpdateOrderRequest {
	return r._tmsToMscUpdateOrderRequest
}

var poolTaobaoLogisticsExpressOrderTmsUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsExpressOrderTmsUpdateRequest()
	},
}

// GetTaobaoLogisticsExpressOrderTmsUpdateRequest 从 sync.Pool 获取 TaobaoLogisticsExpressOrderTmsUpdateAPIRequest
func GetTaobaoLogisticsExpressOrderTmsUpdateAPIRequest() *TaobaoLogisticsExpressOrderTmsUpdateAPIRequest {
	return poolTaobaoLogisticsExpressOrderTmsUpdateAPIRequest.Get().(*TaobaoLogisticsExpressOrderTmsUpdateAPIRequest)
}

// ReleaseTaobaoLogisticsExpressOrderTmsUpdateAPIRequest 将 TaobaoLogisticsExpressOrderTmsUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsExpressOrderTmsUpdateAPIRequest(v *TaobaoLogisticsExpressOrderTmsUpdateAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsExpressOrderTmsUpdateAPIRequest.Put(v)
}
