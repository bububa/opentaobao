package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchannelsupplierproductlist 供应商渠道产品列表查询
// alibaba.ascp.channel.supplier.product.list
//
// 供应商查询渠道产品列表
func Alibabaascpchannelsupplierproductlist(clt *core.SDKClient, req *ascpchannel.AlibabaascpchannelsupplierproductlistAPIRequest, session string) (*ascpchannel.AlibabaascpchannelsupplierproductlistAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchannelsupplierproductlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
