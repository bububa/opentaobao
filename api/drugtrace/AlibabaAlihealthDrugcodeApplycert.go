package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeApplycert 申请证书为对接方
// alibaba.alihealth.drugcode.applycert
//
// 申请证书 为对接方(当前是药厂和中心化系统)
func AlibabaAlihealthDrugcodeApplycert(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeApplycertAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeApplycertAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
