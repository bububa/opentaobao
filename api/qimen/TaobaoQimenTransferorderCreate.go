package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenTransferorderCreate 调拨单创建
// taobao.qimen.transferorder.create
//
// 调拨单创建
func TaobaoQimenTransferorderCreate(clt *core.SDKClient, req *qimen.TaobaoQimenTransferorderCreateAPIRequest, session string) (*qimen.TaobaoQimenTransferorderCreateAPIResponse, error) {
	var resp qimen.TaobaoQimenTransferorderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
