package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRecycleOfnpreredpacketGetAPIRequest 服务商查询前置补贴红包的最新数据 API请求
// taobao.recycle.ofnpreredpacket.get
//
// 服务商查询前置补贴红包的最新数据
type TaobaoRecycleOfnpreredpacketGetAPIRequest struct {
	model.Params
	// 旧机单id
	_oldOrderId int64
}

// NewTaobaoRecycleOfnpreredpacketGetRequest 初始化TaobaoRecycleOfnpreredpacketGetAPIRequest对象
func NewTaobaoRecycleOfnpreredpacketGetRequest() *TaobaoRecycleOfnpreredpacketGetAPIRequest {
	return &TaobaoRecycleOfnpreredpacketGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRecycleOfnpreredpacketGetAPIRequest) GetApiMethodName() string {
	return "taobao.recycle.ofnpreredpacket.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRecycleOfnpreredpacketGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRecycleOfnpreredpacketGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOldOrderId is OldOrderId Setter
// 旧机单id
func (r *TaobaoRecycleOfnpreredpacketGetAPIRequest) SetOldOrderId(_oldOrderId int64) error {
	r._oldOrderId = _oldOrderId
	r.Set("old_order_id", _oldOrderId)
	return nil
}

// GetOldOrderId OldOrderId Getter
func (r TaobaoRecycleOfnpreredpacketGetAPIRequest) GetOldOrderId() int64 {
	return r._oldOrderId
}
