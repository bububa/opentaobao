package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaorecyclepredeductoldget 查询回收单前置抵扣详情
// taobao.recycle.prededuct.old.get
//
// 查询回收单前置抵扣详情
func Taobaorecyclepredeductoldget(clt *core.SDKClient, req *servicecenter.TaobaorecyclepredeductoldgetAPIRequest, session string) (*servicecenter.TaobaorecyclepredeductoldgetAPIResponse, error) {
	var resp servicecenter.TaobaorecyclepredeductoldgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
