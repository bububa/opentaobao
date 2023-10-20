package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpIndustryWaybillPreAccept 商家ERP预推单
// alibaba.ascp.industry.waybill.pre.accept
//
// 商家ERP预推单
func AlibabaAscpIndustryWaybillPreAccept(clt *core.SDKClient, req *ascpchannel.AlibabaAscpIndustryWaybillPreAcceptAPIRequest, session string) (*ascpchannel.AlibabaAscpIndustryWaybillPreAcceptAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpIndustryWaybillPreAcceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
