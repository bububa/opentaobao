package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpChannelDistributorProductDetailAPIRequest
获取供应链渠道中心品的详情接口（淘外分销商专用） API请求
alibaba.ascp.channel.distributor.product.detail

此api为淘外分销的品批量查询标准api，淘外分销商专用 */
type AlibabaAscpChannelDistributorProductDetailAPIRequest struct {
	model.Params
	// 产品详情查询入参
	_productDetailRequest *ProductDetailQueryRequestForDistributor
}

// New
