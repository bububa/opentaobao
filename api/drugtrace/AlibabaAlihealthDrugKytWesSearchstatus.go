package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytwessearchstatus 单据处理状态查询
// alibaba.alihealth.drug.kyt.wes.searchstatus
//
// 单据处理状态查询
func Alibabaalihealthdrugkytwessearchstatus(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytwessearchstatusAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytwessearchstatusAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytwessearchstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
