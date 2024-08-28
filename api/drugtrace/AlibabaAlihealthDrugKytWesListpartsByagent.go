package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesListpartsByagent 物流代货主查找往来单位接口
// alibaba.alihealth.drug.kyt.wes.listparts.byagent
//
// 代理企业查询往来单位列表
func AlibabaAlihealthDrugKytWesListpartsByagent(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesListpartsByagentAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
