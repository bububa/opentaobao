package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcode 根据单据编号查询单据明细
// alibaba.alihealth.drug.kyt.query.specia.vaccin.billcode
//
// 根据单据编号查询单据明细
func AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytQuerySpeciaVaccinBillcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
