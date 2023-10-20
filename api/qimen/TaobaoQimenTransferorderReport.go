package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimentransferorderreport 调拨单通知
// taobao.qimen.transferorder.report
//
// 调拨单通知
func Taobaoqimentransferorderreport(clt *core.SDKClient, req *qimen.TaobaoqimentransferorderreportAPIRequest, session string) (*qimen.TaobaoqimentransferorderreportAPIResponse, error) {
	var resp qimen.TaobaoqimentransferorderreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
