package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaorecyclepredeductblacklistordersync 同步服务商黑名单
// taobao.recycle.prededuct.blacklist.order.sync
//
// 同步服务商黑名单
func Taobaorecyclepredeductblacklistordersync(clt *core.SDKClient, req *servicecenter.TaobaorecyclepredeductblacklistordersyncAPIRequest, session string) (*servicecenter.TaobaorecyclepredeductblacklistordersyncAPIResponse, error) {
	var resp servicecenter.TaobaorecyclepredeductblacklistordersyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
