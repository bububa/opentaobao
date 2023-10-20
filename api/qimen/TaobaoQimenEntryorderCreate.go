package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenEntryorderCreate 入库单创建接口
// taobao.qimen.entryorder.create
//
// taobao.qimen.entryorder.create
func TaobaoQimenEntryorderCreate(clt *core.SDKClient, req *qimen.TaobaoQimenEntryorderCreateAPIRequest, session string) (*qimen.TaobaoQimenEntryorderCreateAPIResponse, error) {
	var resp qimen.TaobaoQimenEntryorderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
