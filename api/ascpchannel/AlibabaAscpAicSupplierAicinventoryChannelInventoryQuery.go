package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpAicSupplierAicinventoryChannelInventoryQuery 供应商直发-商家仓库存查询服务
// alibaba.ascp.aic.supplier.aicinventory.channel.inventory.query
//
// 提供商家基于货品、供应商、仓，查询ascp 实时商家仓库存查询数据。
func AlibabaAscpAicSupplierAicinventoryChannelInventoryQuery(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIRequest, resp *ascpchannel.AlibabaAscpAicSupplierAicinventoryChannelInventoryQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
