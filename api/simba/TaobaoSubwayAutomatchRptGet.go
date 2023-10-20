package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayAutomatchRptGet 查询流量智选天级报告
// taobao.subway.automatch.rpt.get
//
// 查询流量智选天级报告
func TaobaoSubwayAutomatchRptGet(clt *core.SDKClient, req *simba.TaobaoSubwayAutomatchRptGetAPIRequest, resp *simba.TaobaoSubwayAutomatchRptGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
