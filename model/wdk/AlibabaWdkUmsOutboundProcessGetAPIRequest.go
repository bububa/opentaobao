package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkumsoutboundprocessgetAPIRequest 出库业务UMS异步处理结果返回 API请求
// alibaba.wdk.ums.outbound.process.get
//
// 出库业务UMS异步处理结果返回
type AlibabawdkumsoutboundprocessgetAPIRequest struct {
	model.Params
	// 店仓code，指的是作业对象，对应一个物理店或仓编码
	_warehouseCode string
}

// NewAlibabawdkumsoutboundprocessgetRequest 初始化AlibabawdkumsoutboundprocessgetAPIRequest对象
func NewAlibabawdkumsoutboundprocessgetRequest() *AlibabawdkumsoutboundprocessgetAPIRequest {
	return &AlibabawdkumsoutboundprocessgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkumsoutboundprocessgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ums.outbound.process.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkumsoutboundprocessgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkumsoutboundprocessgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCode is WarehouseCode Setter
// 店仓code，指的是作业对象，对应一个物理店或仓编码
func (r *AlibabawdkumsoutboundprocessgetAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r AlibabawdkumsoutboundprocessgetAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}
