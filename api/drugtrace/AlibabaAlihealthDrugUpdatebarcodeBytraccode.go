package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugUpdatebarcodeBytraccode 根据追溯码修改69码
// alibaba.alihealth.drug.updatebarcode.bytraccode
//
// 根据追溯码修改69码
func AlibabaAlihealthDrugUpdatebarcodeBytraccode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
