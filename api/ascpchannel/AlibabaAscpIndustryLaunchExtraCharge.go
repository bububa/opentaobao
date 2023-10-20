package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpIndustryLaunchExtraCharge 阿里巴巴.行业.增加费用.服务商发起
// alibaba.ascp.industry.launch.extra.charge
//
// 阿里巴巴.行业.增加费用.服务商发起
func AlibabaAscpIndustryLaunchExtraCharge(clt *core.SDKClient, req *ascpchannel.AlibabaAscpIndustryLaunchExtraChargeAPIRequest, session string) (*ascpchannel.AlibabaAscpIndustryLaunchExtraChargeAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpIndustryLaunchExtraChargeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
