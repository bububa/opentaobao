package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangConsignorderGiftReverse 赠品绑赠回滚
// alibaba.dchain.aoxiang.consignorder.gift.reverse
//
// 赠品绑赠回滚
func AlibabaDchainAoxiangConsignorderGiftReverse(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest, session string) (*ascp.AlibabaDchainAoxiangConsignorderGiftReverseAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangConsignorderGiftReverseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
