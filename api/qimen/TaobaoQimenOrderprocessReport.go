package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenOrderprocessReport 订单流水通知接口
// taobao.qimen.orderprocess.report
//
// taobao.qimen.orderprocess.report
func TaobaoQimenOrderprocessReport(clt *core.SDKClient, req *qimen.TaobaoQimenOrderprocessReportAPIRequest, session string) (*qimen.TaobaoQimenOrderprocessReportAPIResponse, error) {
	var resp qimen.TaobaoQimenOrderprocessReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
