package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// WdkLogisticNetworkWarehouseDeliveryRelationQuery 仓站（网格仓自提点）关系查询
// wdk.logistic.network.warehouse.delivery.relation.query
//
// 盒马集市，仓站（网格仓自提点）关系查询
func WdkLogisticNetworkWarehouseDeliveryRelationQuery(ctx context.Context, clt *core.SDKClient, req *logistic.WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIRequest, resp *logistic.WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
