package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelSupplierProductDetail 供应链渠道中心分销品详情查询(供应商专用)
// alibaba.ascp.channel.supplier.product.detail
//
// 供应链渠道中心分销品详情查询(供应商专用)
func AlibabaAscpChannelSupplierProductDetail(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelSupplierProductDetailAPIRequest, resp *ascpchannel.AlibabaAscpChannelSupplierProductDetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
