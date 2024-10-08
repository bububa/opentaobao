package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSpeciaVaccinSearchstatus 疫苗企业上传单据后处理状态查询
// alibaba.alihealth.drug.kyt.specia.vaccin.searchstatus
//
// 单据处理状态查询
func AlibabaAlihealthDrugKytSpeciaVaccinSearchstatus(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
