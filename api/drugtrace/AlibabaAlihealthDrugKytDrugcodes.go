package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrugcodes 药品是否赋码
// alibaba.alihealth.drug.kyt.drugcodes
//
// 药品是否赋码
func AlibabaAlihealthDrugKytDrugcodes(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrugcodesAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrugcodesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
