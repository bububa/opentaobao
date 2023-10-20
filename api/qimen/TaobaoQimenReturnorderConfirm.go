package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenReturnorderConfirm 退货入库单确认接口
// taobao.qimen.returnorder.confirm
//
// taobao.qimen.returnorder.confirm
func TaobaoQimenReturnorderConfirm(clt *core.SDKClient, req *qimen.TaobaoQimenReturnorderConfirmAPIRequest, resp *qimen.TaobaoQimenReturnorderConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
