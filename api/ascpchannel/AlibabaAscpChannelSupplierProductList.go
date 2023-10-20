package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelSupplierProductList 供应商渠道产品列表查询
// alibaba.ascp.channel.supplier.product.list
//
// 供应商查询渠道产品列表
func AlibabaAscpChannelSupplierProductList(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelSupplierProductListAPIRequest, session string) (*ascpchannel.AlibabaAscpChannelSupplierProductListAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpChannelSupplierProductListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
