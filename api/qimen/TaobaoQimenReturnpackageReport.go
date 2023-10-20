package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenReturnpackageReport 退货包裹状态通知接口
// taobao.qimen.returnpackage.report
//
// 退货包裹状态通知接口
func TaobaoQimenReturnpackageReport(clt *core.SDKClient, req *qimen.TaobaoQimenReturnpackageReportAPIRequest, session string) (*qimen.TaobaoQimenReturnpackageReportAPIResponse, error) {
	var resp qimen.TaobaoQimenReturnpackageReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
