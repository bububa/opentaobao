package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpChannelDistributorProductListAPIRequest
供应链渠道中心淘外分销品批量查询(分销商专用) API请求
alibaba.ascp.channel.distributor.product.list

此api为淘外分销的品批量查询标准api，淘外分销商专用 */
type AlibabaAscpChannelDistributorProductListAPIRequest struct {
	model.Params
	// 列表请求
	_productListRequest *Productlistrequest
}

// New
