package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytUploadb2cbill 快易通零售B2C
// alibaba.alihealth.drug.kyt.uploadb2cbill
//
// 快易通零售B2C单据上传
func AlibabaAlihealthDrugKytUploadb2cbill(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUploadb2cbillAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytUploadb2cbillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
