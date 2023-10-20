package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenentryordercreate 入库单创建接口
// taobao.qimen.entryorder.create
//
// taobao.qimen.entryorder.create
func Taobaoqimenentryordercreate(clt *core.SDKClient, req *qimen.TaobaoqimenentryordercreateAPIRequest, session string) (*qimen.TaobaoqimenentryordercreateAPIResponse, error) {
	var resp qimen.TaobaoqimenentryordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
