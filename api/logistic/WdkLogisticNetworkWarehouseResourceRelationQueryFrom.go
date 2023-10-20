package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// WdkLogisticNetworkWarehouseResourceRelationQueryFrom 中心仓查网格仓
// wdk.logistic.network.warehouse.resource.relation.query.from
//
// 盒马集市，中心仓查询网格仓
func WdkLogisticNetworkWarehouseResourceRelationQueryFrom(clt *core.SDKClient, req *logistic.WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest, session string) (*logistic.WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse, error) {
	var resp logistic.WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
