package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofulfillmentorderassembleAPIRequest 拆合单结果回传接口 API请求
// taobao.fulfillment.order.assemble
//
// 拆合单结果回传接口
type TaobaofulfillmentorderassembleAPIRequest struct {
	model.Params
	// 操作类型，支持参数为MERGE、CANCEL_MERGE。当进行CANCEL_MERGE操作时，只需要传入groupId即可，order_list可以为空
	_type string
	// 批量回传集合,一次接口最多40
	_assembleOrders *AssembleOrder
}

// NewTaobaofulfillmentorderassembleRequest 初始化TaobaofulfillmentorderassembleAPIRequest对象
func NewTaobaofulfillmentorderassembleRequest() *TaobaofulfillmentorderassembleAPIRequest {
	return &TaobaofulfillmentorderassembleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofulfillmentorderassembleAPIRequest) GetApiMethodName() string {
	return "taobao.fulfillment.order.assemble"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofulfillmentorderassembleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofulfillmentorderassembleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 操作类型，支持参数为MERGE、CANCEL_MERGE。当进行CANCEL_MERGE操作时，只需要传入groupId即可，order_list可以为空
func (r *TaobaofulfillmentorderassembleAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaofulfillmentorderassembleAPIRequest) GetType() string {
	return r._type
}

// SetAssembleOrders is AssembleOrders Setter
// 批量回传集合,一次接口最多40
func (r *TaobaofulfillmentorderassembleAPIRequest) SetAssembleOrders(_assembleOrders *AssembleOrder) error {
	r._assembleOrders = _assembleOrders
	r.Set("assemble_orders", _assembleOrders)
	return nil
}

// GetAssembleOrders AssembleOrders Getter
func (r TaobaofulfillmentorderassembleAPIRequest) GetAssembleOrders() *AssembleOrder {
	return r._assembleOrders
}
