package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSearchstatus 单据处理状态查询
// alibaba.alihealth.drug.kyt.searchstatus
//
// 单据处理状态查询
func AlibabaAlihealthDrugKytSearchstatus(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSearchstatusAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytSearchstatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
