package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpIndustryWaybillAccept 商家ERP预推单
// alibaba.ascp.industry.waybill.accept
//
// 商家ERP预推单
func AlibabaAscpIndustryWaybillAccept(clt *core.SDKClient, req *ascpchannel.AlibabaAscpIndustryWaybillAcceptAPIRequest, session string) (*ascpchannel.AlibabaAscpIndustryWaybillAcceptAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpIndustryWaybillAcceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
