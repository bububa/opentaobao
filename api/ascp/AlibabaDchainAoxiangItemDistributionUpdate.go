package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemDistributionUpdate 更新商品分销内容
// alibaba.dchain.aoxiang.item.distribution.update
//
// 更新商品分销内容
func AlibabaDchainAoxiangItemDistributionUpdate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemDistributionUpdateAPIRequest, session string) (*ascp.AlibabaDchainAoxiangItemDistributionUpdateAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangItemDistributionUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
