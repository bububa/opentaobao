package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugGetbarcodeBytraccode 根据追溯码获取69码
// alibaba.alihealth.drug.getbarcode.bytraccode
//
// 根据追溯码获取69码
func AlibabaAlihealthDrugGetbarcodeBytraccode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
