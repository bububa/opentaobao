package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimentransferordercreate 调拨单创建
// taobao.qimen.transferorder.create
//
// 调拨单创建
func Taobaoqimentransferordercreate(clt *core.SDKClient, req *qimen.TaobaoqimentransferordercreateAPIRequest, session string) (*qimen.TaobaoqimentransferordercreateAPIResponse, error) {
	var resp qimen.TaobaoqimentransferordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
