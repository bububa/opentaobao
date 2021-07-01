package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

/* TaobaoQimenWavenumReport
发货单波次通知接口
taobao.qimen.wavenum.report

WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求 */
func TaobaoQimenWavenumReport(clt *core.SDKClient, req *qimen.TaobaoQimenWavenumReportAPIRequest, session string) (*qimen.TaobaoQimenWavenumReportAPIResponse, error) {
	var resp qimen.TaobaoQimenWavenumReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
