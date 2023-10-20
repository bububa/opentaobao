package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeApplycert 申请证书为对接方
// alibaba.alihealth.drugcode.applycert
//
// 申请证书 为对接方(当前是药厂和中心化系统)
func AlibabaAlihealthDrugcodeApplycert(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeApplycertAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeApplycertAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
