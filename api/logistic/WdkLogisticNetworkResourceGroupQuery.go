package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// WdkLogisticNetworkResourceGroupQuery 查询网格仓-区块-自提点关系
// wdk.logistic.network.resource.group.query
//
// 查询网格仓-区块-自提点关系
func WdkLogisticNetworkResourceGroupQuery(clt *core.SDKClient, req *logistic.WdkLogisticNetworkResourceGroupQueryAPIRequest, session string) (*logistic.WdkLogisticNetworkResourceGroupQueryAPIResponse, error) {
	var resp logistic.WdkLogisticNetworkResourceGroupQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
