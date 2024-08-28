package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytListparts 查询往来单位列表
// alibaba.alihealth.drug.kyt.listparts
//
// 查询往来单位列表
func AlibabaAlihealthDrugKytListparts(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytListpartsAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytListpartsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
