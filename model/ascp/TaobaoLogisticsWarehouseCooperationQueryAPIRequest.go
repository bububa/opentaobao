package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticswarehousecooperationqueryAPIRequest 仓合作关系查询 API请求
// taobao.logistics.warehouse.cooperation.query
//
// 仓合作关系查询
type TaobaologisticswarehousecooperationqueryAPIRequest struct {
	model.Params
	// 仓合作关系查询入参
	_warehouseCooperationQueryRequest *WarehouseCooperationQueryRequest
}

// NewTaobaologisticswarehousecooperationqueryRequest 初始化TaobaologisticswarehousecooperationqueryAPIRequest对象
func NewTaobaologisticswarehousecooperationqueryRequest() *TaobaologisticswarehousecooperationqueryAPIRequest {
	return &TaobaologisticswarehousecooperationqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticswarehousecooperationqueryAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.warehouse.cooperation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticswarehousecooperationqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticswarehousecooperationqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCooperationQueryRequest is WarehouseCooperationQueryRequest Setter
// 仓合作关系查询入参
func (r *TaobaologisticswarehousecooperationqueryAPIRequest) SetWarehouseCooperationQueryRequest(_warehouseCooperationQueryRequest *WarehouseCooperationQueryRequest) error {
	r._warehouseCooperationQueryRequest = _warehouseCooperationQueryRequest
	r.Set("warehouse_cooperation_query_request", _warehouseCooperationQueryRequest)
	return nil
}

// GetWarehouseCooperationQueryRequest WarehouseCooperationQueryRequest Getter
func (r TaobaologisticswarehousecooperationqueryAPIRequest) GetWarehouseCooperationQueryRequest() *WarehouseCooperationQueryRequest {
	return r._warehouseCooperationQueryRequest
}
