package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbownerDelete 民宿房东删除接口
// taobao.xhotel.bnbowner.delete
//
// 民宿房东删除接口，删除房东后，对应的门店及房源会同步删除，请谨慎使用
func TaobaoXhotelBnbownerDelete(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbownerDeleteAPIRequest, resp *xhotelitem.TaobaoXhotelBnbownerDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
