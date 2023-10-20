package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorecycleofnpreredpacketgetAPIRequest 服务商查询前置补贴红包的最新数据 API请求
// taobao.recycle.ofnpreredpacket.get
//
// 服务商查询前置补贴红包的最新数据
type TaobaorecycleofnpreredpacketgetAPIRequest struct {
	model.Params
	// 旧机单id
	_oldOrderId int64
}

// NewTaobaorecycleofnpreredpacketgetRequest 初始化TaobaorecycleofnpreredpacketgetAPIRequest对象
func NewTaobaorecycleofnpreredpacketgetRequest() *TaobaorecycleofnpreredpacketgetAPIRequest {
	return &TaobaorecycleofnpreredpacketgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorecycleofnpreredpacketgetAPIRequest) GetApiMethodName() string {
	return "taobao.recycle.ofnpreredpacket.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorecycleofnpreredpacketgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorecycleofnpreredpacketgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOldOrderId is OldOrderId Setter
// 旧机单id
func (r *TaobaorecycleofnpreredpacketgetAPIRequest) SetOldOrderId(_oldOrderId int64) error {
	r._oldOrderId = _oldOrderId
	r.Set("old_order_id", _oldOrderId)
	return nil
}

// GetOldOrderId OldOrderId Getter
func (r TaobaorecycleofnpreredpacketgetAPIRequest) GetOldOrderId() int64 {
	return r._oldOrderId
}
