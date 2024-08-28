package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesListparts 查询往来单位列表
// alibaba.alihealth.drug.kyt.wes.listparts
//
// 查询往来单位列表
func AlibabaAlihealthDrugKytWesListparts(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesListpartsAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesListpartsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
