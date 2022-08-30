package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemDistributionDescribe 分销商品文描
// alibaba.dchain.aoxiang.item.distribution.describe
//
// 分销商品文描
func AlibabaDchainAoxiangItemDistributionDescribe(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemDistributionDescribeAPIRequest, session string) (*ascp.AlibabaDchainAoxiangItemDistributionDescribeAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangItemDistributionDescribeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
