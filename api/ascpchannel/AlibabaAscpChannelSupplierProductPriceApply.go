package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelSupplierProductPriceApply 供应商设置渠道产品价格
// alibaba.ascp.channel.supplier.product.price.apply
//
// 供应商设置渠道产品价格
func AlibabaAscpChannelSupplierProductPriceApply(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelSupplierProductPriceApplyAPIRequest, session string) (*ascpchannel.AlibabaAscpChannelSupplierProductPriceApplyAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpChannelSupplierProductPriceApplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
