package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenInventorySynchronizeReport 库存状态同步确认接口
// taobao.qimen.inventory.synchronize.report
//
// 库存状态同步确认接口
func TaobaoQimenInventorySynchronizeReport(clt *core.SDKClient, req *qimen.TaobaoQimenInventorySynchronizeReportAPIRequest, session string) (*qimen.TaobaoQimenInventorySynchronizeReportAPIResponse, error) {
	var resp qimen.TaobaoQimenInventorySynchronizeReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
