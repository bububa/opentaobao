package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytsearchstatus 单据处理状态查询
// alibaba.alihealth.drug.kyt.searchstatus
//
// 单据处理状态查询
func Alibabaalihealthdrugkytsearchstatus(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytsearchstatusAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytsearchstatusAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytsearchstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
