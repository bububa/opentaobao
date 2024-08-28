package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangReceiverinfoQuery 供应链优仓即时配隐私小号查询
// alibaba.dchain.aoxiang.receiverinfo.query
//
// 供应链优仓即时配隐私小号查询
func AlibabaDchainAoxiangReceiverinfoQuery(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangReceiverinfoQueryAPIRequest, resp *ascp.AlibabaDchainAoxiangReceiverinfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
