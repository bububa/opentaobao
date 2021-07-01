package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpChannelDistributorProductSelectAPIRequest
供应链渠道中心品的选品接口（淘外分销商专用） API请求
alibaba.ascp.channel.distributor.product.select

此api为淘外分销的品的选品标准api，淘外分销商专用 */
type AlibabaAscpChannelDistributorProductSelectAPIRequest struct {
	model.Params
	// 选品请求
	_selectProductRequest *ProductLinkRequest
}

// New
