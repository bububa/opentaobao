package alihealthcrm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaFmhealthWeightLossplanSynclossplan 减重计划--同步减重计划
// alibaba.fmhealth.weight.lossplan.synclossplan
//
// 减重计划--三方同步用户初始化减重计划给我们
func AlibabaFmhealthWeightLossplanSynclossplan(ctx context.Context, clt *core.SDKClient, req *alihealthcrm.AlibabaFmhealthWeightLossplanSynclossplanAPIRequest, resp *alihealthcrm.AlibabaFmhealthWeightLossplanSynclossplanAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
