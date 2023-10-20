package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchannelsupplierproductauth 供应商授权渠道产品到市场或分销商
// alibaba.ascp.channel.supplier.product.auth
//
// 供应商授权渠道产品到市场或分销商
func Alibabaascpchannelsupplierproductauth(clt *core.SDKClient, req *ascpchannel.AlibabaascpchannelsupplierproductauthAPIRequest, session string) (*ascpchannel.AlibabaascpchannelsupplierproductauthAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchannelsupplierproductauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
