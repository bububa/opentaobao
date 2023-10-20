package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugBillUpbillDetailWithcode 查询上游出库单明细(带追溯码信息)
// alibaba.alihealth.drug.bill.upbill.detail.withcode
//
// 查询上游出库单明细(带追溯码信息)
func AlibabaAlihealthDrugBillUpbillDetailWithcode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugBillUpbillDetailWithcodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugBillUpbillDetailWithcodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
