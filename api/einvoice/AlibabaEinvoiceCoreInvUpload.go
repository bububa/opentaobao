package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceCoreInvUpload 发票中台-发票结果回传
// alibaba.einvoice.core.inv.upload
//
// 发票回传接口适用于以下场景：
// ① 阿里发票平台向ISV提交原始发票申请，ISV开具发票成功后，基于申请ID(apply_id)回传发票至阿里发票平台进行归集与交付。
// ② 直接回传发票给阿里发票平台，进行归集，并交付给业务前台和用户。
func AlibabaEinvoiceCoreInvUpload(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceCoreInvUploadAPIRequest, resp *einvoice.AlibabaEinvoiceCoreInvUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
