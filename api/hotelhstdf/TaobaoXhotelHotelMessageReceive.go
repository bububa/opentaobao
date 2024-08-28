package hotelhstdf

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// TaobaoXhotelHotelMessageReceive 接收道消息接口
// taobao.xhotel.hotel.message.receive
//
// 接收道消息接口
func TaobaoXhotelHotelMessageReceive(ctx context.Context, clt *core.SDKClient, req *hotelhstdf.TaobaoXhotelHotelMessageReceiveAPIRequest, resp *hotelhstdf.TaobaoXhotelHotelMessageReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
