package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugWxinfoUpload 小程序数据回传
// alibaba.alihealth.drug.wxinfo.upload
//
// 小程序数据回传
func AlibabaAlihealthDrugWxinfoUpload(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugWxinfoUploadAPIRequest, resp *drugtrace.AlibabaAlihealthDrugWxinfoUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
