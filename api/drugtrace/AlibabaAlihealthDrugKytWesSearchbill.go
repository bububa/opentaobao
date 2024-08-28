package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesSearchbill 通过时间段批量查询入出库单信息
// alibaba.alihealth.drug.kyt.wes.searchbill
//
// 通过时间段批量查询入出库单信息
func AlibabaAlihealthDrugKytWesSearchbill(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesSearchbillAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesSearchbillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
