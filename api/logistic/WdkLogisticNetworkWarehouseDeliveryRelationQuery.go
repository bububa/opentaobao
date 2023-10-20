package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Wdklogisticnetworkwarehousedeliveryrelationquery 仓站（网格仓自提点）关系查询
// wdk.logistic.network.warehouse.delivery.relation.query
//
// 盒马集市，仓站（网格仓自提点）关系查询
func Wdklogisticnetworkwarehousedeliveryrelationquery(clt *core.SDKClient, req *logistic.WdklogisticnetworkwarehousedeliveryrelationqueryAPIRequest, session string) (*logistic.WdklogisticnetworkwarehousedeliveryrelationqueryAPIResponse, error) {
	var resp logistic.WdklogisticnetworkwarehousedeliveryrelationqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
