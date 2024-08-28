package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesListupout 查询货主/本企业上游企业出库单据信息
// alibaba.alihealth.drug.kyt.wes.listupout
//
// 查询货主/本企业上游企业出库单据信息
func AlibabaAlihealthDrugKytWesListupout(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesListupoutAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesListupoutAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
