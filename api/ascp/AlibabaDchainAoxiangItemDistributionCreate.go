package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemDistributionCreate 选择店铺商品并进行铺货（通用）
// alibaba.dchain.aoxiang.item.distribution.create
//
// 选择店铺商品并进行铺货, 铺货给所有的合作分销商。设定的价格为通用价格
func AlibabaDchainAoxiangItemDistributionCreate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemDistributionCreateAPIRequest, session string) (*ascp.AlibabaDchainAoxiangItemDistributionCreateAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangItemDistributionCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
