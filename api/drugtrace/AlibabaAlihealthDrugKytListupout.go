package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytListupout 查询货主/本企业上游企业出库单据信息
// alibaba.alihealth.drug.kyt.listupout
//
// 查询货主/本企业上游企业出库单据信息
func AlibabaAlihealthDrugKytListupout(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytListupoutAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytListupoutAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
