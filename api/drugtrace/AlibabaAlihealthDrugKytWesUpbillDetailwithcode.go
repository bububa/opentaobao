package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesUpbillDetailwithcode 查询上游出库单明细（带追溯码信息）
// alibaba.alihealth.drug.kyt.wes.upbill.detailwithcode
//
// 查询上游出库单明细(带追溯码信息)
func AlibabaAlihealthDrugKytWesUpbillDetailwithcode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytWesUpbillDetailwithcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
