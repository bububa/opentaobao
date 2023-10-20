package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenreturnordercreate 退货入库单创建接口
// taobao.qimen.returnorder.create
//
// taobao.qimen.returnorder.create
func Taobaoqimenreturnordercreate(clt *core.SDKClient, req *qimen.TaobaoqimenreturnordercreateAPIRequest, session string) (*qimen.TaobaoqimenreturnordercreateAPIResponse, error) {
	var resp qimen.TaobaoqimenreturnordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
