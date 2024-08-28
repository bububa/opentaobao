package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqySearchstatus 单据处理状态查询
// alibaba.alihealth.drug.kyt.scqy.searchstatus
//
// 单据处理状态查询
func AlibabaAlihealthDrugKytScqySearchstatus(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqySearchstatusAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytScqySearchstatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
