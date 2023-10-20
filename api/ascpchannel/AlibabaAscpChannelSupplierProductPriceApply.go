package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchannelsupplierproductpriceapply 供应商设置渠道产品价格
// alibaba.ascp.channel.supplier.product.price.apply
//
// 供应商设置渠道产品价格
func Alibabaascpchannelsupplierproductpriceapply(clt *core.SDKClient, req *ascpchannel.AlibabaascpchannelsupplierproductpriceapplyAPIRequest, session string) (*ascpchannel.AlibabaascpchannelsupplierproductpriceapplyAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchannelsupplierproductpriceapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
