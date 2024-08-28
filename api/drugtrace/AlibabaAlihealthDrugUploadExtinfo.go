package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugUploadExtinfo 中药饮片及器械对接
// alibaba.alihealth.drug.upload.extinfo
//
// 中药饮片及器械对接
func AlibabaAlihealthDrugUploadExtinfo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugUploadExtinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugUploadExtinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
