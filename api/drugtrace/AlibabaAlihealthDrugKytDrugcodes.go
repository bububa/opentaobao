package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrugcodes 药品是否赋码
// alibaba.alihealth.drug.kyt.drugcodes
//
// 药品是否赋码
func AlibabaAlihealthDrugKytDrugcodes(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrugcodesAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrugcodesAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
