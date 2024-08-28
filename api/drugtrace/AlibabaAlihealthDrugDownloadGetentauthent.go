package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugDownloadGetentauthent 获取授权企业列表
// alibaba.alihealth.drug.download.getentauthent
//
// D2D数据落地获取授权企业列表
func AlibabaAlihealthDrugDownloadGetentauthent(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugDownloadGetentauthentAPIRequest, resp *drugtrace.AlibabaAlihealthDrugDownloadGetentauthentAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
