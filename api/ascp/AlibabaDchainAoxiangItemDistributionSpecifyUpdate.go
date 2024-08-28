package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemDistributionSpecifyUpdate 指定分销商进行铺货(专享) - 修改
// alibaba.dchain.aoxiang.item.distribution.specify.update
//
// 选择店铺的商品进行指定分销商铺货。 可以指定对应的分销商对应的价格
func AlibabaDchainAoxiangItemDistributionSpecifyUpdate(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIRequest, resp *ascp.AlibabaDchainAoxiangItemDistributionSpecifyUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
