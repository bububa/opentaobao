package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesSearchstatus 单据处理状态查询
// alibaba.alihealth.drug.kyt.wes.searchstatus
//
// 单据处理状态查询
func AlibabaAlihealthDrugKytWesSearchstatus(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesSearchstatusAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesSearchstatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
