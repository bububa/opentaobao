package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpInventorySyncAPIRequest
库存同步接口 API请求
alibaba.ascp.inventory.sync

新零售联盟商家库存同步接口，用于商家商品库存数量同步到阿里系统 */
type AlibabaAscpInventorySyncAPIRequest struct {
	model.Params
	// 库存同步DTO
	_invSingleItemSyncDto *InvSingleItemSyncDto
}

// New
