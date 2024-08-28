package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytListauths 企业搜索自己授权的物流企业
// alibaba.alihealth.drug.kyt.listauths
//
// 企业搜索自己授权的物流企业
func AlibabaAlihealthDrugKytListauths(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytListauthsAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytListauthsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
