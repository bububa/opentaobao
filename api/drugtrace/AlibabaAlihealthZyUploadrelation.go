package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthZyUploadrelation 中药片关联关系上传
// alibaba.alihealth.zy.uploadrelation
//
// 中药片关联关系上传
func AlibabaAlihealthZyUploadrelation(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthZyUploadrelationAPIRequest, resp *drugtrace.AlibabaAlihealthZyUploadrelationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
