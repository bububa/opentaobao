package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Wdklogisticnetworkwarehouseresourcerelationqueryfrom 中心仓查网格仓
// wdk.logistic.network.warehouse.resource.relation.query.from
//
// 盒马集市，中心仓查询网格仓
func Wdklogisticnetworkwarehouseresourcerelationqueryfrom(clt *core.SDKClient, req *logistic.WdklogisticnetworkwarehouseresourcerelationqueryfromAPIRequest, session string) (*logistic.WdklogisticnetworkwarehouseresourcerelationqueryfromAPIResponse, error) {
	var resp logistic.WdklogisticnetworkwarehouseresourcerelationqueryfromAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
