package perfect

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatcwmsoutboundpickreceiveAPIRequest 拣货接单 API请求
// alibaba.tcwms.outbound.pick.receive
//
// 拣货接单
type AlibabatcwmsoutboundpickreceiveAPIRequest struct {
	model.Params
	// 入口参数
	_pickReceiveRequest *PickReceiveRequest
}

// NewAlibabatcwmsoutboundpickreceiveRequest 初始化AlibabatcwmsoutboundpickreceiveAPIRequest对象
func NewAlibabatcwmsoutboundpickreceiveRequest() *AlibabatcwmsoutboundpickreceiveAPIRequest {
	return &AlibabatcwmsoutboundpickreceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatcwmsoutboundpickreceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.tcwms.outbound.pick.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatcwmsoutboundpickreceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatcwmsoutboundpickreceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPickReceiveRequest is PickReceiveRequest Setter
// 入口参数
func (r *AlibabatcwmsoutboundpickreceiveAPIRequest) SetPickReceiveRequest(_pickReceiveRequest *PickReceiveRequest) error {
	r._pickReceiveRequest = _pickReceiveRequest
	r.Set("pick_receive_request", _pickReceiveRequest)
	return nil
}

// GetPickReceiveRequest PickReceiveRequest Getter
func (r AlibabatcwmsoutboundpickreceiveAPIRequest) GetPickReceiveRequest() *PickReceiveRequest {
	return r._pickReceiveRequest
}
