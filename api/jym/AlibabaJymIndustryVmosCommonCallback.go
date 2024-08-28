package jym

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymIndustryVmosCommonCallback vmos游戏信息采集结果回调通知
// alibaba.jym.industry.vmos.common.callback
//
// vmos游戏信息采集结果回调通知
func AlibabaJymIndustryVmosCommonCallback(ctx context.Context, clt *core.SDKClient, req *jym.AlibabaJymIndustryVmosCommonCallbackAPIRequest, resp *jym.AlibabaJymIndustryVmosCommonCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
