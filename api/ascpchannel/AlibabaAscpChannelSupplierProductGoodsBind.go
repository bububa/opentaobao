package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchannelsupplierproductgoodsbind 渠道产品与货品绑定接口
// alibaba.ascp.channel.supplier.product.goods.bind
//
// 渠道产品与货品绑定接口
func Alibabaascpchannelsupplierproductgoodsbind(clt *core.SDKClient, req *ascpchannel.AlibabaascpchannelsupplierproductgoodsbindAPIRequest, session string) (*ascpchannel.AlibabaascpchannelsupplierproductgoodsbindAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchannelsupplierproductgoodsbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
