package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmshopvipcancel 卖家取消店铺vip的优惠
// taobao.crm.shopvip.cancel
//
// 此接口用于取消VIP优惠
func Taobaocrmshopvipcancel(clt *core.SDKClient, req *crm.TaobaocrmshopvipcancelAPIRequest, session string) (*crm.TaobaocrmshopvipcancelAPIResponse, error) {
	var resp crm.TaobaocrmshopvipcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
