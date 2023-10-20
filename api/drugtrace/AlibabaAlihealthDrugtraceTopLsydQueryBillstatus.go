package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetoplsydquerybillstatus 上传单据后处理状态查询
// alibaba.alihealth.drugtrace.top.lsyd.query.billstatus
//
// 单据处理状态查询
func Alibabaalihealthdrugtracetoplsydquerybillstatus(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetoplsydquerybillstatusAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetoplsydquerybillstatusAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetoplsydquerybillstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
