package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Taobaoplacestoreextendadd 新增门店扩展属性
// taobao.place.store.extend.add
//
// 新增授权用户的门店扩展属性
func Taobaoplacestoreextendadd(clt *core.SDKClient, req *alsc.TaobaoplacestoreextendaddAPIRequest, session string) (*alsc.TaobaoplacestoreextendaddAPIResponse, error) {
	var resp alsc.TaobaoplacestoreextendaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
