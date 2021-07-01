package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugGetbarcodeBytraccode
根据追溯码获取69码
alibaba.alihealth.drug.getbarcode.bytraccode

根据追溯码获取69码 */
func AlibabaAlihealthDrugGetbarcodeBytraccode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
