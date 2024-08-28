package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytRemnantbillUpload 零头出入库单据上传
// alibaba.alihealth.drug.kyt.remnantbill.upload
//
// 零头出入库单据上传
func AlibabaAlihealthDrugKytRemnantbillUpload(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytRemnantbillUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
