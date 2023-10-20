package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Taobaoplacestorecategoryget 获取门店类目信息
// taobao.place.storecategory.get
//
// 获取门店类目信息
func Taobaoplacestorecategoryget(clt *core.SDKClient, req *alsc.TaobaoplacestorecategorygetAPIRequest, session string) (*alsc.TaobaoplacestorecategorygetAPIResponse, error) {
	var resp alsc.TaobaoplacestorecategorygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
