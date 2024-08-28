package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpGlobalSupplierItemListInfoQuery 供应商供品信息查询
// alibaba.ascp.global.supplier.item.list.info.query
//
// 供应商供品信息查询
func AlibabaAscpGlobalSupplierItemListInfoQuery(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest, resp *ascpchannel.AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
