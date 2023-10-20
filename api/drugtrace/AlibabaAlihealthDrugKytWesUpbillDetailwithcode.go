package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesUpbillDetailwithcode 查询上游出库单明细（带追溯码信息）
// alibaba.alihealth.drug.kyt.wes.upbill.detailwithcode
//
// 查询上游出库单明细(带追溯码信息)
func AlibabaAlihealthDrugKytWesUpbillDetailwithcode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
