package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSearchstatus 单据处理状态查询
// alibaba.alihealth.drug.kyt.searchstatus
//
// 单据处理状态查询
func AlibabaAlihealthDrugKytSearchstatus(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSearchstatusAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytSearchstatusAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytSearchstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
