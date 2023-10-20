package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchannelsupplierproductdetail 供应链渠道中心分销品详情查询(供应商专用)
// alibaba.ascp.channel.supplier.product.detail
//
// 供应链渠道中心分销品详情查询(供应商专用)
func Alibabaascpchannelsupplierproductdetail(clt *core.SDKClient, req *ascpchannel.AlibabaascpchannelsupplierproductdetailAPIRequest, session string) (*ascpchannel.AlibabaascpchannelsupplierproductdetailAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchannelsupplierproductdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
