package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenstoredelete 门店删除接口
// taobao.qimen.store.delete
//
// 商家在ERP等系统中调用该接口，删除线下门店
func Taobaoqimenstoredelete(clt *core.SDKClient, req *qimen.TaobaoqimenstoredeleteAPIRequest, session string) (*qimen.TaobaoqimenstoredeleteAPIResponse, error) {
	var resp qimen.TaobaoqimenstoredeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
