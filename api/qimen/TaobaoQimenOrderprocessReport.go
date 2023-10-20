package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenorderprocessreport 订单流水通知接口
// taobao.qimen.orderprocess.report
//
// taobao.qimen.orderprocess.report
func Taobaoqimenorderprocessreport(clt *core.SDKClient, req *qimen.TaobaoqimenorderprocessreportAPIRequest, session string) (*qimen.TaobaoqimenorderprocessreportAPIResponse, error) {
	var resp qimen.TaobaoqimenorderprocessreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
