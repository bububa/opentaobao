package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqySearchstatus 单据处理状态查询
// alibaba.alihealth.drug.kyt.scqy.searchstatus
//
// 单据处理状态查询
func AlibabaAlihealthDrugKytScqySearchstatus(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqySearchstatusAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytScqySearchstatusAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytScqySearchstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
