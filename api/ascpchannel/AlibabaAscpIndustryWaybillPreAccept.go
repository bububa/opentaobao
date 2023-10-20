package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpindustrywaybillpreaccept 商家ERP预推单
// alibaba.ascp.industry.waybill.pre.accept
//
// 商家ERP预推单
func Alibabaascpindustrywaybillpreaccept(clt *core.SDKClient, req *ascpchannel.AlibabaascpindustrywaybillpreacceptAPIRequest, session string) (*ascpchannel.AlibabaascpindustrywaybillpreacceptAPIResponse, error) {
	var resp ascpchannel.AlibabaascpindustrywaybillpreacceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
