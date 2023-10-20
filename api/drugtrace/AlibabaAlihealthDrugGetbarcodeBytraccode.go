package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdruggetbarcodebytraccode 根据追溯码获取69码
// alibaba.alihealth.drug.getbarcode.bytraccode
//
// 根据追溯码获取69码
func Alibabaalihealthdruggetbarcodebytraccode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdruggetbarcodebytraccodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdruggetbarcodebytraccodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdruggetbarcodebytraccodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
