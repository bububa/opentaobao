package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbpromoDelete 民宿卖家活动删除
// taobao.xhotel.bnbpromo.delete
//
// 民宿删除营销活动
func TaobaoXhotelBnbpromoDelete(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbpromoDeleteAPIRequest, resp *xhotelitem.TaobaoXhotelBnbpromoDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
