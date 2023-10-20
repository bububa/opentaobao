package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimentransferorderquery 调拨单查询
// taobao.qimen.transferorder.query
//
// 调拨单查询
func Taobaoqimentransferorderquery(clt *core.SDKClient, req *qimen.TaobaoqimentransferorderqueryAPIRequest, session string) (*qimen.TaobaoqimentransferorderqueryAPIResponse, error) {
	var resp qimen.TaobaoqimentransferorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
