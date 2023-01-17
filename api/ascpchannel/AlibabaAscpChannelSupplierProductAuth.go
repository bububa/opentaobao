package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelSupplierProductAuth 供应商授权渠道产品到市场或分销商
// alibaba.ascp.channel.supplier.product.auth
//
// 供应商授权渠道产品到市场或分销商
func AlibabaAscpChannelSupplierProductAuth(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelSupplierProductAuthAPIRequest, session string) (*ascpchannel.AlibabaAscpChannelSupplierProductAuthAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpChannelSupplierProductAuthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
