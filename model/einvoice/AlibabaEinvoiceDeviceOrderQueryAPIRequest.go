package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceDeviceOrderQueryAPIRequest 查询税控设备加盘订购单详情 API请求
// alibaba.einvoice.device.order.query
//
// 查询税控设备订购单详情
type AlibabaEinvoiceDeviceOrderQueryAPIRequest struct {
	model.Params
	// 税控设备订购单ID
	_flowId string
}

// NewAlibabaEinvoiceDeviceOrderQueryRequest 初始化AlibabaEinvoiceDeviceOrderQueryAPIRequest对象
func NewAlibabaEinvoiceDeviceOrderQueryRequest() *AlibabaEinvoiceDeviceOrderQueryAPIRequest {
	return &AlibabaEinvoiceDeviceOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceDeviceOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.device.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceDeviceOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceDeviceOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFlowId is FlowId Setter
// 税控设备订购单ID
func (r *AlibabaEinvoiceDeviceOrderQueryAPIRequest) SetFlowId(_flowId string) error {
	r._flowId = _flowId
	r.Set("flow_id", _flowId)
	return nil
}

// GetFlowId FlowId Getter
func (r AlibabaEinvoiceDeviceOrderQueryAPIRequest) GetFlowId() string {
	return r._flowId
}
