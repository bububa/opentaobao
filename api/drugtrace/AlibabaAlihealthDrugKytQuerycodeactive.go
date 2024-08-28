package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytQuerycodeactive 查询码是否激活
// alibaba.alihealth.drug.kyt.querycodeactive
//
// 查询码是否激活
func AlibabaAlihealthDrugKytQuerycodeactive(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytQuerycodeactiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
