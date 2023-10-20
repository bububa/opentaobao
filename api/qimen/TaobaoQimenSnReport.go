package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimensnreport 发货单SN通知接口
// taobao.qimen.sn.report
//
// WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP
func Taobaoqimensnreport(clt *core.SDKClient, req *qimen.TaobaoqimensnreportAPIRequest, session string) (*qimen.TaobaoqimensnreportAPIResponse, error) {
	var resp qimen.TaobaoqimensnreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
