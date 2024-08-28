package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesQuerycodeactive 查询码是否激活
// alibaba.alihealth.drug.kyt.wes.querycodeactive
//
// 查询码是否激活
func AlibabaAlihealthDrugKytWesQuerycodeactive(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesQuerycodeactiveAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
