package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpIndustryLaunchExtraCharge 阿里巴巴.行业.增加费用.服务商发起
// alibaba.ascp.industry.launch.extra.charge
//
// 阿里巴巴.行业.增加费用.服务商发起
func AlibabaAscpIndustryLaunchExtraCharge(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpIndustryLaunchExtraChargeAPIRequest, resp *ascpchannel.AlibabaAscpIndustryLaunchExtraChargeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
