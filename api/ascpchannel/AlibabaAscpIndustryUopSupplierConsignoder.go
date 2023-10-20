package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpindustryuopsupplierconsignoder 商家推单
// alibaba.ascp.industry.uop.supplier.consignoder
//
// 商家推单
func Alibabaascpindustryuopsupplierconsignoder(clt *core.SDKClient, req *ascpchannel.AlibabaascpindustryuopsupplierconsignoderAPIRequest, session string) (*ascpchannel.AlibabaascpindustryuopsupplierconsignoderAPIResponse, error) {
	var resp ascpchannel.AlibabaascpindustryuopsupplierconsignoderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
