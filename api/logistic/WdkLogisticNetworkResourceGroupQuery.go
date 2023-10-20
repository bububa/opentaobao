package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Wdklogisticnetworkresourcegroupquery 查询网格仓-区块-自提点关系
// wdk.logistic.network.resource.group.query
//
// 查询网格仓-区块-自提点关系
func Wdklogisticnetworkresourcegroupquery(clt *core.SDKClient, req *logistic.WdklogisticnetworkresourcegroupqueryAPIRequest, session string) (*logistic.WdklogisticnetworkresourcegroupqueryAPIResponse, error) {
	var resp logistic.WdklogisticnetworkresourcegroupqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
