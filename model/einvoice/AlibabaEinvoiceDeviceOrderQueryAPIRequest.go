package einvoice

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceDeviceOrderQueryAPIRequest) Reset() {
	r._flowId = ""
	r.Params.ToZero()
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

var poolAlibabaEinvoiceDeviceOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceDeviceOrderQueryRequest()
	},
}

// GetAlibabaEinvoiceDeviceOrderQueryRequest 从 sync.Pool 获取 AlibabaEinvoiceDeviceOrderQueryAPIRequest
func GetAlibabaEinvoiceDeviceOrderQueryAPIRequest() *AlibabaEinvoiceDeviceOrderQueryAPIRequest {
	return poolAlibabaEinvoiceDeviceOrderQueryAPIRequest.Get().(*AlibabaEinvoiceDeviceOrderQueryAPIRequest)
}

// ReleaseAlibabaEinvoiceDeviceOrderQueryAPIRequest 将 AlibabaEinvoiceDeviceOrderQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceDeviceOrderQueryAPIRequest(v *AlibabaEinvoiceDeviceOrderQueryAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceDeviceOrderQueryAPIRequest.Put(v)
}
