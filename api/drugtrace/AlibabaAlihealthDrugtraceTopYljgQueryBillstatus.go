package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetopyljgquerybillstatus 上传单据后处理状态查询
// alibaba.alihealth.drugtrace.top.yljg.query.billstatus
//
// 单据处理状态查询
func Alibabaalihealthdrugtracetopyljgquerybillstatus(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetopyljgquerybillstatusAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetopyljgquerybillstatusAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetopyljgquerybillstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
