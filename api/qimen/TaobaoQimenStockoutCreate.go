package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenstockoutcreate 出库单创建接口
// taobao.qimen.stockout.create
//
// taobao.qimen.returnorder.create
func Taobaoqimenstockoutcreate(clt *core.SDKClient, req *qimen.TaobaoqimenstockoutcreateAPIRequest, session string) (*qimen.TaobaoqimenstockoutcreateAPIResponse, error) {
	var resp qimen.TaobaoqimenstockoutcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
