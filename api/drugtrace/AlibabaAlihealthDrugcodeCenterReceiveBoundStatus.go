package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodecenterreceiveboundstatus 接收中央随机化系统和临床研究机构的绑定确认状态
// alibaba.alihealth.drugcode.center.receive.bound.status
//
// 临床用药试验-接收中央随机化系统和临床研究机构的绑定确认状态
func Alibabaalihealthdrugcodecenterreceiveboundstatus(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodecenterreceiveboundstatusAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodecenterreceiveboundstatusAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodecenterreceiveboundstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
