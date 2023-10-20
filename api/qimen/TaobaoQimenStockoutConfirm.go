package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenstockoutconfirm 出库单确认接口
// taobao.qimen.stockout.confirm
//
// 货品出库后，WMS将状态回传给ERP
func Taobaoqimenstockoutconfirm(clt *core.SDKClient, req *qimen.TaobaoqimenstockoutconfirmAPIRequest, session string) (*qimen.TaobaoqimenstockoutconfirmAPIResponse, error) {
	var resp qimen.TaobaoqimenstockoutconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
