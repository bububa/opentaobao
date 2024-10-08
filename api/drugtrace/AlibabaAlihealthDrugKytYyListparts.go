package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytYyListparts 查询往来单位
// alibaba.alihealth.drug.kyt.yy.listparts
//
// 查询往来单位列表
func AlibabaAlihealthDrugKytYyListparts(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytYyListpartsAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytYyListpartsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
