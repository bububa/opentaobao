package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytFiledownload 处理失败单据下载
// alibaba.alihealth.drug.kyt.filedownload
//
// 处理失败单据下载
func AlibabaAlihealthDrugKytFiledownload(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytFiledownloadAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytFiledownloadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
