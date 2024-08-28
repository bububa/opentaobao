package hotel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotel"
)

// TaobaoXhotelInfoListGet 酒店详细信息查询
// taobao.xhotel.info.list.get
//
// 获取酒店详情信息
func TaobaoXhotelInfoListGet(ctx context.Context, clt *core.SDKClient, req *hotel.TaobaoXhotelInfoListGetAPIRequest, resp *hotel.TaobaoXhotelInfoListGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
