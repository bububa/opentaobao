package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpChannelDistributorInventoryGetAPIRequest
链渠道中心淘外库存查询 API请求
alibaba.ascp.channel.distributor.inventory.get

此api为淘外分销的渠道产品库存查询标准api，淘外分销商专用 */
type AlibabaAscpChannelDistributorInventoryGetAPIRequest struct {
	model.Params
	// 入参
	_inventoryRequest *ChannelInventoryQuery
}

// New
