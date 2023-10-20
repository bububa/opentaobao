package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangreceiverinfoquery 供应链优仓即时配隐私小号查询
// alibaba.dchain.aoxiang.receiverinfo.query
//
// 供应链优仓即时配隐私小号查询
func Alibabadchainaoxiangreceiverinfoquery(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangreceiverinfoqueryAPIRequest, session string) (*ascp.AlibabadchainaoxiangreceiverinfoqueryAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangreceiverinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
