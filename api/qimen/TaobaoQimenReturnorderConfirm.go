package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenReturnorderConfirm 退货入库单确认接口
// taobao.qimen.returnorder.confirm
//
// taobao.qimen.returnorder.confirm
func TaobaoQimenReturnorderConfirm(clt *core.SDKClient, req *qimen.TaobaoQimenReturnorderConfirmAPIRequest, session string) (*qimen.TaobaoQimenReturnorderConfirmAPIResponse, error) {
	var resp qimen.TaobaoQimenReturnorderConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
