package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelSupplierProductList 供应商渠道产品列表查询
// alibaba.ascp.channel.supplier.product.list
//
// 供应商查询渠道产品列表
func AlibabaAscpChannelSupplierProductList(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelSupplierProductListAPIRequest, resp *ascpchannel.AlibabaAscpChannelSupplierProductListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
