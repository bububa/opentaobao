package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDruploadretail 快易通多融零售上传接口
// alibaba.alihealth.drug.kyt.druploadretail
//
// 快易通多融零售上传接口
func AlibabaAlihealthDrugKytDruploadretail(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDruploadretailAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDruploadretailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
