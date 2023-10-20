package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenwavenumreport 发货单波次通知接口
// taobao.qimen.wavenum.report
//
// WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求
func Taobaoqimenwavenumreport(clt *core.SDKClient, req *qimen.TaobaoqimenwavenumreportAPIRequest, session string) (*qimen.TaobaoqimenwavenumreportAPIResponse, error) {
	var resp qimen.TaobaoqimenwavenumreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
