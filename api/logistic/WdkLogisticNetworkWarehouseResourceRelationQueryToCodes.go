package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// WdkLogisticNetworkWarehouseResourceRelationQueryToCodes 按网格仓查中心仓（带缓存）
// wdk.logistic.network.warehouse.resource.relation.query.to.codes
//
// 盒马集市，网格仓查询中心仓
func WdkLogisticNetworkWarehouseResourceRelationQueryToCodes(clt *core.SDKClient, req *logistic.WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest, session string) (*logistic.WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse, error) {
	var resp logistic.WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
