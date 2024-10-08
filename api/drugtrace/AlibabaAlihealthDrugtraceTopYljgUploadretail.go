package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgUploadretail 零售单据上传接口
// alibaba.alihealth.drugtrace.top.yljg.uploadretail
//
// 医疗机构零售上传接口
func AlibabaAlihealthDrugtraceTopYljgUploadretail(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgUploadretailAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgUploadretailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
