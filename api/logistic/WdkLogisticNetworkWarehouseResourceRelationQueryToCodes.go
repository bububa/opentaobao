package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// WdkLogisticNetworkWarehouseResourceRelationQueryToCodes 按网格仓查中心仓（带缓存）
// wdk.logistic.network.warehouse.resource.relation.query.to.codes
//
// 盒马集市，网格仓查询中心仓
func WdkLogisticNetworkWarehouseResourceRelationQueryToCodes(clt *core.SDKClient, req *logistic.WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest, resp *logistic.WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
