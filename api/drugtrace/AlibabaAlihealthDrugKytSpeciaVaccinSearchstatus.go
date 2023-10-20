package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSpeciaVaccinSearchstatus 疫苗企业上传单据后处理状态查询
// alibaba.alihealth.drug.kyt.specia.vaccin.searchstatus
//
// 单据处理状态查询
func AlibabaAlihealthDrugKytSpeciaVaccinSearchstatus(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
