package icburfq

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// AlibabaIcbuAnnexUpload 上传附件获取附件files_str
// alibaba.icbu.annex.upload
//
// 上传附件获取附件files_str
func AlibabaIcbuAnnexUpload(ctx context.Context, clt *core.SDKClient, req *icburfq.AlibabaIcbuAnnexUploadAPIRequest, resp *icburfq.AlibabaIcbuAnnexUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
