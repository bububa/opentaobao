package alihealthmedical

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmedical"
)

// AlibabaAlihealthMedicalImDataUpload 三方IM图片音频消息上传
// alibaba.alihealth.medical.im.data.upload
//
// 三方IM图片音频消息上传
func AlibabaAlihealthMedicalImDataUpload(ctx context.Context, clt *core.SDKClient, req *alihealthmedical.AlibabaAlihealthMedicalImDataUploadAPIRequest, resp *alihealthmedical.AlibabaAlihealthMedicalImDataUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
