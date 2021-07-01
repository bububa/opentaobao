package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpChannelDistributorInventoryListGetAPIRequest
批量查询渠道库存 API请求
alibaba.ascp.channel.distributor.inventory.list.get

淘外分销批量查询渠道产品的库存 */
type AlibabaAscpChannelDistributorInventoryListGetAPIRequest struct {
	model.Params
	// 系统自动生成
	_inventoryRequest *BatchChannelInventoryQuery
}

// New
