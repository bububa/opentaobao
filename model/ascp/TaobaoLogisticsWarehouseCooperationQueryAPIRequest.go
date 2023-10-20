package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWarehouseCooperationQueryAPIRequest 仓合作关系查询 API请求
// taobao.logistics.warehouse.cooperation.query
//
// 仓合作关系查询
type TaobaoLogisticsWarehouseCooperationQueryAPIRequest struct {
	model.Params
	// 仓合作关系查询入参
	_warehouseCooperationQueryRequest *WarehouseCooperationQueryRequest
}

// NewTaobaoLogisticsWarehouseCooperationQueryRequest 初始化TaobaoLogisticsWarehouseCooperationQueryAPIRequest对象
func NewTaobaoLogisticsWarehouseCooperationQueryRequest() *TaobaoLogisticsWarehouseCooperationQueryAPIRequest {
	return &TaobaoLogisticsWarehouseCooperationQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsWarehouseCooperationQueryAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.warehouse.cooperation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsWarehouseCooperationQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsWarehouseCooperationQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWarehouseCooperationQueryRequest is WarehouseCooperationQueryRequest Setter
// 仓合作关系查询入参
func (r *TaobaoLogisticsWarehouseCooperationQueryAPIRequest) SetWarehouseCooperationQueryRequest(_warehouseCooperationQueryRequest *WarehouseCooperationQueryRequest) error {
	r._warehouseCooperationQueryRequest = _warehouseCooperationQueryRequest
	r.Set("warehouse_cooperation_query_request", _warehouseCooperationQueryRequest)
	return nil
}

// GetWarehouseCooperationQueryRequest WarehouseCooperationQueryRequest Getter
func (r TaobaoLogisticsWarehouseCooperationQueryAPIRequest) GetWarehouseCooperationQueryRequest() *WarehouseCooperationQueryRequest {
	return r._warehouseCooperationQueryRequest
}
