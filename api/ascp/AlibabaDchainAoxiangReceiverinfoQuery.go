package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangReceiverinfoQuery 供应链优仓即时配隐私小号查询
// alibaba.dchain.aoxiang.receiverinfo.query
//
// 供应链优仓即时配隐私小号查询
func AlibabaDchainAoxiangReceiverinfoQuery(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangReceiverinfoQueryAPIRequest, session string) (*ascp.AlibabaDchainAoxiangReceiverinfoQueryAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangReceiverinfoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
