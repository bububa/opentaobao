package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemDistributionDescribe 分销商品文描
// alibaba.dchain.aoxiang.item.distribution.describe
//
// 分销商品文描
func AlibabaDchainAoxiangItemDistributionDescribe(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemDistributionDescribeAPIRequest, resp *ascp.AlibabaDchainAoxiangItemDistributionDescribeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
