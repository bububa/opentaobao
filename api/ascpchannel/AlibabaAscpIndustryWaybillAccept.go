package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpindustrywaybillaccept 商家ERP预推单
// alibaba.ascp.industry.waybill.accept
//
// 商家ERP预推单
func Alibabaascpindustrywaybillaccept(clt *core.SDKClient, req *ascpchannel.AlibabaascpindustrywaybillacceptAPIRequest, session string) (*ascpchannel.AlibabaascpindustrywaybillacceptAPIResponse, error) {
	var resp ascpchannel.AlibabaascpindustrywaybillacceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
