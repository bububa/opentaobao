package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytQueryDruginfoFromBillcode 根据单据编号查询单据明细
// alibaba.alihealth.drug.kyt.query.druginfo.from.billcode
//
// 根据单据编号查询单据明细
func AlibabaAlihealthDrugKytQueryDruginfoFromBillcode(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
