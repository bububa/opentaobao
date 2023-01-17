package rhino

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRhinoSupplychainOutboundPickingcompleteAPIRequest 【WMS005】接收成衣捡配完成通知 API请求
// taobao.rhino.supplychain.outbound.pickingcomplete
//
// 接收成衣捡配完成通知,WMS005
type TaobaoRhinoSupplychainOutboundPickingcompleteAPIRequest struct {
	model.Params
	// 捡配完成消息
	_param0 *PickingCompleteMsg
}

// NewTaobaoRhinoSupplychainOutboundPickingcompleteRequest 初始化TaobaoRhinoSupplychainOutboundPickingcompleteAPIRequest对象
func NewTaobaoRhinoSupplychainOutboundPickingcompleteRequest() *TaobaoRhinoSupplychainOutboundPickingcompleteAPIRequest {
	return &TaobaoRhinoSupplychainOutboundPickingcompleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRhinoSupplychainOutboundPickingcompleteAPIRequest) GetApiMethodName() string {
	return "taobao.rhino.supplychain.outbound.pickingcomplete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRhinoSupplychainOutboundPickingcompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRhinoSupplychainOutboundPickingcompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 捡配完成消息
func (r *TaobaoRhinoSupplychainOutboundPickingcompleteAPIRequest) SetParam0(_param0 *PickingCompleteMsg) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoRhinoSupplychainOutboundPickingcompleteAPIRequest) GetParam0() *PickingCompleteMsg {
	return r._param0
}
