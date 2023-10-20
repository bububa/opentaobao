package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenordersnreport 订单SN通知接口
// taobao.qimen.order.sn.report
//
// WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表
func Taobaoqimenordersnreport(clt *core.SDKClient, req *qimen.TaobaoqimenordersnreportAPIRequest, session string) (*qimen.TaobaoqimenordersnreportAPIResponse, error) {
	var resp qimen.TaobaoqimenordersnreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
