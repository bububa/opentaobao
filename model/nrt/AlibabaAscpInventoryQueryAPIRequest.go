package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpInventoryQueryAPIRequest
商品库存查询接口 API请求
alibaba.ascp.inventory.query

新零售联盟商家库存查询接口，用于商家商品库存数量感知查询 */
type AlibabaAscpInventoryQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_invSingleItemSyncDto *InvSingleItemSyncDto
}

// New
