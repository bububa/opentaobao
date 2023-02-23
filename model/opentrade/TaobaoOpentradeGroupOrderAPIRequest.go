package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeGroupOrderAPIRequest 组团购获取订单列表 API请求
// taobao.opentrade.group.order
//
// 组团购场景下，获取开团的订单列表
type TaobaoOpentradeGroupOrderAPIRequest struct {
	model.Params
	// 团id
	_groupId int64
}

// NewTaobaoOpentradeGroupOrderRequest 初始化TaobaoOpentradeGroupOrderAPIRequest对象
func NewTaobaoOpentradeGroupOrderRequest() *TaobaoOpentradeGroupOrderAPIRequest {
	return &TaobaoOpentradeGroupOrderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupOrderAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.group.order"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupOrderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpentradeGroupOrderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupId is GroupId Setter
// 团id
func (r *TaobaoOpentradeGroupOrderAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r TaobaoOpentradeGroupOrderAPIRequest) GetGroupId() int64 {
	return r._groupId
}
