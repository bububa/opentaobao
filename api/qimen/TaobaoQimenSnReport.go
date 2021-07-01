package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

/* TaobaoQimenSnReport
发货单SN通知接口
taobao.qimen.sn.report

WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP */
func TaobaoQimenSnReport(clt *core.SDKClient, req *qimen.TaobaoQimenSnReportAPIRequest, session string) (*qimen.TaobaoQimenSnReportAPIResponse, error) {
	var resp qimen.TaobaoQimenSnReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
