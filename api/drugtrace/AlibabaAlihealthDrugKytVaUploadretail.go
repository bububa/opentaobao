package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytVaUploadretail 零售单据上传接口
// alibaba.alihealth.drug.kyt.va.uploadretail
//
// 零售上传单据信息接口
func AlibabaAlihealthDrugKytVaUploadretail(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytVaUploadretailAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytVaUploadretailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
