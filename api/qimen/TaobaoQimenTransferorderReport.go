package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenTransferorderReport 调拨单通知
// taobao.qimen.transferorder.report
//
// 调拨单通知
func TaobaoQimenTransferorderReport(clt *core.SDKClient, req *qimen.TaobaoQimenTransferorderReportAPIRequest, resp *qimen.TaobaoQimenTransferorderReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
