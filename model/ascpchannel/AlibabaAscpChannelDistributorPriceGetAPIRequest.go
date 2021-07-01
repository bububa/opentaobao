package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpChannelDistributorPriceGetAPIRequest
链渠道中心淘外分销价格查询(分销商专用) API请求
alibaba.ascp.channel.distributor.price.get

此api为淘外分销的渠道产品价格查询标准api，淘外分销商专用 */
type AlibabaAscpChannelDistributorPriceGetAPIRequest struct {
	model.Params
	// 价格入参
	_priceRequest *Pricerequest
}

// New
