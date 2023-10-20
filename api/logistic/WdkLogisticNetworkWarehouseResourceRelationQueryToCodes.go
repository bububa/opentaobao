package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Wdklogisticnetworkwarehouseresourcerelationquerytocodes 按网格仓查中心仓（带缓存）
// wdk.logistic.network.warehouse.resource.relation.query.to.codes
//
// 盒马集市，网格仓查询中心仓
func Wdklogisticnetworkwarehouseresourcerelationquerytocodes(clt *core.SDKClient, req *logistic.WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIRequest, session string) (*logistic.WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIResponse, error) {
	var resp logistic.WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
