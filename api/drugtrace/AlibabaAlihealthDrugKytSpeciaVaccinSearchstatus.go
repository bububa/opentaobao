package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSpeciaVaccinSearchstatus 疫苗企业上传单据后处理状态查询
// alibaba.alihealth.drug.kyt.specia.vaccin.searchstatus
//
// 单据处理状态查询
func AlibabaAlihealthDrugKytSpeciaVaccinSearchstatus(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
