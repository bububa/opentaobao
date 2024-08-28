package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcode 根据单据编号查询单据明细
// alibaba.alihealth.drug.kyt.query.specia.vaccin.billcode
//
// 根据单据编号查询单据明细
func AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcode(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
