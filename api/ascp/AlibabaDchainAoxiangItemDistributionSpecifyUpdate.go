package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemDistributionSpecifyUpdate 指定分销商进行铺货(专享) - 修改
// alibaba.dchain.aoxiang.item.distribution.specify.update
//
// 选择店铺的商品进行指定分销商铺货。 可以指定对应的分销商对应的价格
func AlibabaDchainAoxiangItemDistributionSpecifyUpdate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest, session string) (*ascp.AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
