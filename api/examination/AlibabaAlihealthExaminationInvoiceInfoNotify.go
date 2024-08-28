package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationInvoiceInfoNotify 体检机构同步发票信息给阿里健康
// alibaba.alihealth.examination.invoice.info.notify
//
// 体检机构向阿里健康同步发票信息
func AlibabaAlihealthExaminationInvoiceInfoNotify(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationInvoiceInfoNotifyAPIRequest, resp *examination.AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
