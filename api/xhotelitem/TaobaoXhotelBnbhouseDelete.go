package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbhouseDelete 民宿门店删除接口
// taobao.xhotel.bnbhouse.delete
//
// 支持门店相关的门店删除，删除门店会级联删除门店下面的房源
func TaobaoXhotelBnbhouseDelete(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbhouseDeleteAPIRequest, resp *xhotelitem.TaobaoXhotelBnbhouseDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
