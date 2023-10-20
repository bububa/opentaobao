package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugupdatebarcodebytraccode 根据追溯码修改69码
// alibaba.alihealth.drug.updatebarcode.bytraccode
//
// 根据追溯码修改69码
func Alibabaalihealthdrugupdatebarcodebytraccode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugupdatebarcodebytraccodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugupdatebarcodebytraccodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugupdatebarcodebytraccodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
