package topoaid

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFulfillmentOrderAssembleAPIRequest 拆合单结果回传接口 API请求
// taobao.fulfillment.order.assemble
//
// 拆合单结果回传接口
type TaobaoFulfillmentOrderAssembleAPIRequest struct {
	model.Params
	// 操作类型，支持参数为MERGE、CANCEL_MERGE。当进行CANCEL_MERGE操作时，只需要传入groupId即可，order_list可以为空
	_type string
	// 批量回传集合,一次接口最多40
	_assembleOrders *AssembleOrder
}

// NewTaobaoFulfillmentOrderAssembleRequest 初始化TaobaoFulfillmentOrderAssembleAPIRequest对象
func NewTaobaoFulfillmentOrderAssembleRequest() *TaobaoFulfillmentOrderAssembleAPIRequest {
	return &TaobaoFulfillmentOrderAssembleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFulfillmentOrderAssembleAPIRequest) GetApiMethodName() string {
	return "taobao.fulfillment.order.assemble"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFulfillmentOrderAssembleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetType is Type Setter
// 操作类型，支持参数为MERGE、CANCEL_MERGE。当进行CANCEL_MERGE操作时，只需要传入groupId即可，order_list可以为空
func (r *TaobaoFulfillmentOrderAssembleAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoFulfillmentOrderAssembleAPIRequest) GetType() string {
	return r._type
}

// SetAssembleOrders is AssembleOrders Setter
// 批量回传集合,一次接口最多40
func (r *TaobaoFulfillmentOrderAssembleAPIRequest) SetAssembleOrders(_assembleOrders *AssembleOrder) error {
	r._assembleOrders = _assembleOrders
	r.Set("assemble_orders", _assembleOrders)
	return nil
}

// GetAssembleOrders AssembleOrders Getter
func (r TaobaoFulfillmentOrderAssembleAPIRequest) GetAssembleOrders() *AssembleOrder {
	return r._assembleOrders
}
