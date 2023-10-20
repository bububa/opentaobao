package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenOrderprocessReport 订单流水通知接口
// taobao.qimen.orderprocess.report
//
// taobao.qimen.orderprocess.report
func TaobaoQimenOrderprocessReport(clt *core.SDKClient, req *qimen.TaobaoQimenOrderprocessReportAPIRequest, resp *qimen.TaobaoQimenOrderprocessReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
