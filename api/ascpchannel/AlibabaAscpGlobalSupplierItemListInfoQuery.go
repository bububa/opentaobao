package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpGlobalSupplierItemListInfoQuery 供应商供品信息查询
// alibaba.ascp.global.supplier.item.list.info.query
//
// 供应商供品信息查询
func AlibabaAscpGlobalSupplierItemListInfoQuery(clt *core.SDKClient, req *ascpchannel.AlibabaAscpGlobalSupplierItemListInfoQueryAPIRequest, resp *ascpchannel.AlibabaAscpGlobalSupplierItemListInfoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
