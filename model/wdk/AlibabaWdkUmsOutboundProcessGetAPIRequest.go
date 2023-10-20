package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsOutboundProcessGetAPIRequest 出库业务UMS异步处理结果返回 API请求
// alibaba.wdk.ums.outbound.process.get
//
// 出库业务UMS异步处理结果返回
type AlibabaWdkUmsOutboundProcessGetAPIRequest struct {
	model.Params
	// 店仓code，指的是作业对象，对应一个物理店或仓编码
	_warehouseCode string
}

// NewAlibabaWdkUmsOutboundProcessGetRequest 初始化AlibabaWdkUmsOutboundProcessGetAPIRequest对象
func NewAlibabaWdkUmsOutboundProcessGetRequest() *AlibabaWdkUmsOutboundProcessGetAPIRequest {
	return &AlibabaWdkUmsOutboundProcessGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkUmsOutboundProcessGetAPIRequest) Reset() {
	r._warehouseCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsOutboundProcessGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.outbound.process.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsOutboundProcessGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkUmsOutboundProcessGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 店仓code，指的是作业对象，对应一个物理店或仓编码
func (r *AlibabaWdkUmsOutboundProcessGetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabaWdkUmsOutboundProcessGetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

var poolAlibabaWdkUmsOutboundProcessGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkUmsOutboundProcessGetRequest()
	},
}

// GetAlibabaWdkUmsOutboundProcessGetRequest 从 sync.Pool 获取 AlibabaWdkUmsOutboundProcessGetAPIRequest
func GetAlibabaWdkUmsOutboundProcessGetAPIRequest() *AlibabaWdkUmsOutboundProcessGetAPIRequest {
	return poolAlibabaWdkUmsOutboundProcessGetAPIRequest.Get().(*AlibabaWdkUmsOutboundProcessGetAPIRequest)
}

// ReleaseAlibabaWdkUmsOutboundProcessGetAPIRequest 将 AlibabaWdkUmsOutboundProcessGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkUmsOutboundProcessGetAPIRequest(v *AlibabaWdkUmsOutboundProcessGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkUmsOutboundProcessGetAPIRequest.Put(v)
}
