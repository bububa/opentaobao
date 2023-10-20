package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangReceiverinfoQuery 供应链优仓即时配隐私小号查询
// alibaba.dchain.aoxiang.receiverinfo.query
//
// 供应链优仓即时配隐私小号查询
func AlibabaDchainAoxiangReceiverinfoQuery(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangReceiverinfoQueryAPIRequest, resp *ascp.AlibabaDchainAoxiangReceiverinfoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
