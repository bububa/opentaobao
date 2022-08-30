package perfect

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTcwmsOutboundLoadContainerReceiveAPIRequest 装箱接单 API请求
// alibaba.tcwms.outbound.load.container.receive
//
// 装箱接单
type AlibabaTcwmsOutboundLoadContainerReceiveAPIRequest struct {
	model.Params
	// 参数
	_loadReceiveRequest *LoadReceiveRequest
}

// NewAlibabaTcwmsOutboundLoadContainerReceiveRequest 初始化AlibabaTcwmsOutboundLoadContainerReceiveAPIRequest对象
func NewAlibabaTcwmsOutboundLoadContainerReceiveRequest() *AlibabaTcwmsOutboundLoadContainerReceiveAPIRequest {
	return &AlibabaTcwmsOutboundLoadContainerReceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTcwmsOutboundLoadContainerReceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.tcwms.outbound.load.container.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTcwmsOutboundLoadContainerReceiveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLoadReceiveRequest is LoadReceiveRequest Setter
// 参数
func (r *AlibabaTcwmsOutboundLoadContainerReceiveAPIRequest) SetLoadReceiveRequest(_loadReceiveRequest *LoadReceiveRequest) error {
	r._loadReceiveRequest = _loadReceiveRequest
	r.Set("load_receive_request", _loadReceiveRequest)
	return nil
}

// GetLoadReceiveRequest LoadReceiveRequest Getter
func (r AlibabaTcwmsOutboundLoadContainerReceiveAPIRequest) GetLoadReceiveRequest() *LoadReceiveRequest {
	return r._loadReceiveRequest
}
