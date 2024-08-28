package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// WdkLogisticNetworkResourceGroupQuery 查询网格仓-区块-自提点关系
// wdk.logistic.network.resource.group.query
//
// 查询网格仓-区块-自提点关系
func WdkLogisticNetworkResourceGroupQuery(ctx context.Context, clt *core.SDKClient, req *logistic.WdkLogisticNetworkResourceGroupQueryAPIRequest, resp *logistic.WdkLogisticNetworkResourceGroupQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
