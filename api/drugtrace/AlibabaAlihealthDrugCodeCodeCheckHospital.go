package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeCodeCheckHospital 码核查状态同步-医院
// alibaba.alihealth.drug.code.code.check.hospital
//
// 码核查状态同步-医院
func AlibabaAlihealthDrugCodeCodeCheckHospital(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeCodeCheckHospitalAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugCodeCodeCheckHospitalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
