package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Taobaoplacestoremodify 商家修改线下门店
// taobao.place.store.modify
//
// 用于商家修改线下门店信息
func Taobaoplacestoremodify(clt *core.SDKClient, req *alsc.TaobaoplacestoremodifyAPIRequest, session string) (*alsc.TaobaoplacestoremodifyAPIResponse, error) {
	var resp alsc.TaobaoplacestoremodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
