package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelSupplierProductDetail 供应链渠道中心分销品详情查询(供应商专用)
// alibaba.ascp.channel.supplier.product.detail
//
// 供应链渠道中心分销品详情查询(供应商专用)
func AlibabaAscpChannelSupplierProductDetail(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelSupplierProductDetailAPIRequest, session string) (*ascpchannel.AlibabaAscpChannelSupplierProductDetailAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpChannelSupplierProductDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
