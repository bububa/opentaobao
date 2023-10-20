package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenreturnorderconfirm 退货入库单确认接口
// taobao.qimen.returnorder.confirm
//
// taobao.qimen.returnorder.confirm
func Taobaoqimenreturnorderconfirm(clt *core.SDKClient, req *qimen.TaobaoqimenreturnorderconfirmAPIRequest, session string) (*qimen.TaobaoqimenreturnorderconfirmAPIResponse, error) {
	var resp qimen.TaobaoqimenreturnorderconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
