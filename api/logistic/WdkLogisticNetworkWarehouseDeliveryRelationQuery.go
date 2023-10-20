package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// WdkLogisticNetworkWarehouseDeliveryRelationQuery 仓站（网格仓自提点）关系查询
// wdk.logistic.network.warehouse.delivery.relation.query
//
// 盒马集市，仓站（网格仓自提点）关系查询
func WdkLogisticNetworkWarehouseDeliveryRelationQuery(clt *core.SDKClient, req *logistic.WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest, resp *logistic.WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
