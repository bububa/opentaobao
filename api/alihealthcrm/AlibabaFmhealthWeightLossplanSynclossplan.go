package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaFmhealthWeightLossplanSynclossplan 减重计划--同步减重计划
// alibaba.fmhealth.weight.lossplan.synclossplan
//
// 减重计划--三方同步用户初始化减重计划给我们
func AlibabaFmhealthWeightLossplanSynclossplan(clt *core.SDKClient, req *alihealthcrm.AlibabaFmhealthWeightLossplanSynclossplanAPIRequest, session string) (*alihealthcrm.AlibabaFmhealthWeightLossplanSynclossplanAPIResponse, error) {
	var resp alihealthcrm.AlibabaFmhealthWeightLossplanSynclossplanAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
