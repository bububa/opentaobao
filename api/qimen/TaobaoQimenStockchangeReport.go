package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenstockchangereport 库存异动通知接口
// taobao.qimen.stockchange.report
//
// taobao.qimen.stockchange.report
func Taobaoqimenstockchangereport(clt *core.SDKClient, req *qimen.TaobaoqimenstockchangereportAPIRequest, session string) (*qimen.TaobaoqimenstockchangereportAPIResponse, error) {
	var resp qimen.TaobaoqimenstockchangereportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
