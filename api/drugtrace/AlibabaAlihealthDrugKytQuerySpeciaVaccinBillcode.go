package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytqueryspeciavaccinbillcode 根据单据编号查询单据明细
// alibaba.alihealth.drug.kyt.query.specia.vaccin.billcode
//
// 根据单据编号查询单据明细
func Alibabaalihealthdrugkytqueryspeciavaccinbillcode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytqueryspeciavaccinbillcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
