package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest
供应商直发-商家仓库存查询服务 API请求
alibaba.ascp.aic.supplier.aicinventory.channel.inventory.query

提供商家基于货品、供应商、仓，查询ascp 实时商家仓库存查询数据。 */
type AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest struct {
	model.Params
	// 商家仓库存查询请求参数
	_merchantInventoryQueryRequest *MerchantInventoryQuery
}

// New
