package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodeplatformCodeEntscan 药品商家扫码
// alibaba.alihealth.tracecodeplatform.code.entscan
//
// 药品商家扫描药品监管码，只有该商家的药才返回
func AlibabaAlihealthTracecodeplatformCodeEntscan(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest, resp *alihealth2.AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
