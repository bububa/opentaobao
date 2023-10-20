package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenentryorderquery 入库单查询接口
// taobao.qimen.entryorder.query
//
// ERP调用接口，查询入库单信息;
func Taobaoqimenentryorderquery(clt *core.SDKClient, req *qimen.TaobaoqimenentryorderqueryAPIRequest, session string) (*qimen.TaobaoqimenentryorderqueryAPIResponse, error) {
	var resp qimen.TaobaoqimenentryorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
