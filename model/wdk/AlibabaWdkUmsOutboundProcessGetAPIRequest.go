package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsOutboundProcessGetAPIRequest
出库业务UMS异步处理结果返回 API请求
alibaba.wdk.ums.outbound.process.get

出库业务UMS异步处理结果返回 */
type AlibabaWdkUmsOutboundProcessGetAPIRequest struct {
	model.Params
	// 店仓code，指的是作业对象，对应一个物理店或仓编码
	_warehouseCode string
}

// NewAlibabaWdkUmsOutboundProcessGetRequest 初始化AlibabaWdkUmsOutboundProcessGetAPIRequest对象
func NewAlibabaWdkUmsOutboundProcessGetRequest() *AlibabaWdkUmsOutboundProcessGetAPIRequest {
	return &AlibabaWdkUmsOutboundProcessGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsOutboundProcessGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.outbound.process.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsOutboundProcessGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WarehouseCode Setter
// 店仓code，指的是作业对象，对应一个物理店或仓编码
func (r *AlibabaWdkUmsOutboundProcessGetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// Get WarehouseCode Getter
func (r AlibabaWdkUmsOutboundProcessGetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}
