package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenentryorderconfirm 入库单确认接口
// taobao.qimen.entryorder.confirm
//
// WMS调用接口，回传入库单信息;
func Taobaoqimenentryorderconfirm(clt *core.SDKClient, req *qimen.TaobaoqimenentryorderconfirmAPIRequest, session string) (*qimen.TaobaoqimenentryorderconfirmAPIResponse, error) {
	var resp qimen.TaobaoqimenentryorderconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
