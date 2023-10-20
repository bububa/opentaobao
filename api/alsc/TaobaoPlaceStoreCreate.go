package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Taobaoplacestorecreate 商户门店创建接口
// taobao.place.store.create
//
// 用于商家创建线下门店
func Taobaoplacestorecreate(clt *core.SDKClient, req *alsc.TaobaoplacestorecreateAPIRequest, session string) (*alsc.TaobaoplacestorecreateAPIResponse, error) {
	var resp alsc.TaobaoplacestorecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
