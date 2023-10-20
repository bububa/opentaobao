package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenStockchangeReport 库存异动通知接口
// taobao.qimen.stockchange.report
//
// taobao.qimen.stockchange.report
func TaobaoQimenStockchangeReport(clt *core.SDKClient, req *qimen.TaobaoQimenStockchangeReportAPIRequest, session string) (*qimen.TaobaoQimenStockchangeReportAPIResponse, error) {
	var resp qimen.TaobaoQimenStockchangeReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
