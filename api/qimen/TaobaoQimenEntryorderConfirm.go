package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenEntryorderConfirm 入库单确认接口
// taobao.qimen.entryorder.confirm
//
// WMS调用接口，回传入库单信息;
func TaobaoQimenEntryorderConfirm(clt *core.SDKClient, req *qimen.TaobaoQimenEntryorderConfirmAPIRequest, resp *qimen.TaobaoQimenEntryorderConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
