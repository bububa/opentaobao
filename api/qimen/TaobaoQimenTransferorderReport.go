package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenTransferorderReport 调拨单通知
// taobao.qimen.transferorder.report
//
// 调拨单通知
func TaobaoQimenTransferorderReport(clt *core.SDKClient, req *qimen.TaobaoQimenTransferorderReportAPIRequest, session string) (*qimen.TaobaoQimenTransferorderReportAPIResponse, error) {
	var resp qimen.TaobaoQimenTransferorderReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
