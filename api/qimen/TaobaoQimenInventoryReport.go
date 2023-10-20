package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimeninventoryreport 库存盘点通知接口
// taobao.qimen.inventory.report
//
// WMS调用奇门的接口,将库存盘点情况回传ERP
func Taobaoqimeninventoryreport(clt *core.SDKClient, req *qimen.TaobaoqimeninventoryreportAPIRequest, session string) (*qimen.TaobaoqimeninventoryreportAPIResponse, error) {
	var resp qimen.TaobaoqimeninventoryreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
