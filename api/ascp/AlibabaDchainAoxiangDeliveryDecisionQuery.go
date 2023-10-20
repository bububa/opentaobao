package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangdeliverydecisionquery 查询黑白名单快递
// alibaba.dchain.aoxiang.delivery.decision.query
//
// 查询黑白名单快递
func Alibabadchainaoxiangdeliverydecisionquery(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangdeliverydecisionqueryAPIRequest, session string) (*ascp.AlibabadchainaoxiangdeliverydecisionqueryAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangdeliverydecisionqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
