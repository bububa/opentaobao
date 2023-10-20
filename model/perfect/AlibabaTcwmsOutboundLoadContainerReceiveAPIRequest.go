package perfect

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatcwmsoutboundloadcontainerreceiveAPIRequest 装箱接单 API请求
// alibaba.tcwms.outbound.load.container.receive
//
// 装箱接单
type AlibabatcwmsoutboundloadcontainerreceiveAPIRequest struct {
	model.Params
	// 参数
	_loadReceiveRequest *LoadReceiveRequest
}

// NewAlibabatcwmsoutboundloadcontainerreceiveRequest 初始化AlibabatcwmsoutboundloadcontainerreceiveAPIRequest对象
func NewAlibabatcwmsoutboundloadcontainerreceiveRequest() *AlibabatcwmsoutboundloadcontainerreceiveAPIRequest {
	return &AlibabatcwmsoutboundloadcontainerreceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatcwmsoutboundloadcontainerreceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.tcwms.outbound.load.container.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatcwmsoutboundloadcontainerreceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatcwmsoutboundloadcontainerreceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLoadReceiveRequest is LoadReceiveRequest Setter
// 参数
func (r *AlibabatcwmsoutboundloadcontainerreceiveAPIRequest) SetLoadReceiveRequest(_loadReceiveRequest *LoadReceiveRequest) error {
	r._loadReceiveRequest = _loadReceiveRequest
	r.Set("load_receive_request", _loadReceiveRequest)
	return nil
}

// GetLoadReceiveRequest LoadReceiveRequest Getter
func (r AlibabatcwmsoutboundloadcontainerreceiveAPIRequest) GetLoadReceiveRequest() *LoadReceiveRequest {
	return r._loadReceiveRequest
}
