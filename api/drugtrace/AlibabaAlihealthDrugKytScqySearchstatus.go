package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytscqysearchstatus 单据处理状态查询
// alibaba.alihealth.drug.kyt.scqy.searchstatus
//
// 单据处理状态查询
func Alibabaalihealthdrugkytscqysearchstatus(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytscqysearchstatusAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytscqysearchstatusAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytscqysearchstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
