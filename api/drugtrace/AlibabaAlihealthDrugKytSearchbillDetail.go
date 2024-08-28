package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSearchbillDetail 查询单据详情
// alibaba.alihealth.drug.kyt.searchbill.detail
//
// 根据单据号码查询码单据详情和码信息
func AlibabaAlihealthDrugKytSearchbillDetail(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSearchbillDetailAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytSearchbillDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
