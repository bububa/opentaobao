package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemDistributionCreate 选择店铺商品并进行铺货（通用）
// alibaba.dchain.aoxiang.item.distribution.create
//
// 选择店铺商品并进行铺货, 铺货给所有的合作分销商。设定的价格为通用价格
func AlibabaDchainAoxiangItemDistributionCreate(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemDistributionCreateAPIRequest, resp *ascp.AlibabaDchainAoxiangItemDistributionCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
