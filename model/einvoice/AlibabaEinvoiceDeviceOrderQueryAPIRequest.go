package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicedeviceorderqueryAPIRequest 查询税控设备加盘订购单详情 API请求
// alibaba.einvoice.device.order.query
//
// 查询税控设备订购单详情
type AlibabaeinvoicedeviceorderqueryAPIRequest struct {
	model.Params
	// 税控设备订购单ID
	_flowId string
}

// NewAlibabaeinvoicedeviceorderqueryRequest 初始化AlibabaeinvoicedeviceorderqueryAPIRequest对象
func NewAlibabaeinvoicedeviceorderqueryRequest() *AlibabaeinvoicedeviceorderqueryAPIRequest {
	return &AlibabaeinvoicedeviceorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicedeviceorderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.device.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicedeviceorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicedeviceorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFlowId is FlowId Setter
// 税控设备订购单ID
func (r *AlibabaeinvoicedeviceorderqueryAPIRequest) SetFlowId(_flowId string) error {
	r._flowId = _flowId
	r.Set("flow_id", _flowId)
	return nil
}

// GetFlowId FlowId Getter
func (r AlibabaeinvoicedeviceorderqueryAPIRequest) GetFlowId() string {
	return r._flowId
}
