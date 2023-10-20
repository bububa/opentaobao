package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest 快递送货上门能力同步/更新接口 API请求
// taobao.logistics.express.delivery.send.ability.async
//
// 快递送货上门能力同步/更新接口
type TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest struct {
	model.Params
	// 快递送货上门能力
	_deliverySendAbilityRequest *DeliverySendAbilityRequest
}

// NewTaobaoLogisticsExpressDeliverySendAbilityAsyncRequest 初始化TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest对象
func NewTaobaoLogisticsExpressDeliverySendAbilityAsyncRequest() *TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest {
	return &TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest) Reset() {
	r._deliverySendAbilityRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.delivery.send.ability.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliverySendAbilityRequest is DeliverySendAbilityRequest Setter
// 快递送货上门能力
func (r *TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest) SetDeliverySendAbilityRequest(_deliverySendAbilityRequest *DeliverySendAbilityRequest) error {
	r._deliverySendAbilityRequest = _deliverySendAbilityRequest
	r.Set("delivery_send_ability_request", _deliverySendAbilityRequest)
	return nil
}

// GetDeliverySendAbilityRequest DeliverySendAbilityRequest Getter
func (r TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest) GetDeliverySendAbilityRequest() *DeliverySendAbilityRequest {
	return r._deliverySendAbilityRequest
}

var poolTaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsExpressDeliverySendAbilityAsyncRequest()
	},
}

// GetTaobaoLogisticsExpressDeliverySendAbilityAsyncRequest 从 sync.Pool 获取 TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest
func GetTaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest() *TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest {
	return poolTaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest.Get().(*TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest)
}

// ReleaseTaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest 将 TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest(v *TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsExpressDeliverySendAbilityAsyncAPIRequest.Put(v)
}
