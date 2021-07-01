package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest
仓站（网格仓自提点）关系查询 API请求
wdk.logistic.network.warehouse.delivery.relation.query

盒马集市，仓站（网格仓自提点）关系查询 */
type WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest struct {
	model.Params
	// 参数
	_paramWarehouseDeliveryRelationPageQueryRequest *WarehouseDeliveryRelationPageQueryRequest
}

// New
