package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Taobaoplacestoredelete 线下门店删除
// taobao.place.store.delete
//
// 用于商家删除线下门店
func Taobaoplacestoredelete(clt *core.SDKClient, req *alsc.TaobaoplacestoredeleteAPIRequest, session string) (*alsc.TaobaoplacestoredeleteAPIResponse, error) {
	var resp alsc.TaobaoplacestoredeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
