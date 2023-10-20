package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrSearchstatus 疫苗企业上传单据后处理状态查询
// alibaba.alihealth.drug.kyt.dr.searchstatus
//
// 单据处理状态查询
func AlibabaAlihealthDrugKytDrSearchstatus(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrSearchstatusAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrSearchstatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
