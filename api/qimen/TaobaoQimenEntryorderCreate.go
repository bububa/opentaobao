package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenEntryorderCreate 入库单创建接口
// taobao.qimen.entryorder.create
//
// taobao.qimen.entryorder.create
func TaobaoQimenEntryorderCreate(clt *core.SDKClient, req *qimen.TaobaoQimenEntryorderCreateAPIRequest, resp *qimen.TaobaoQimenEntryorderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
