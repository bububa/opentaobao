package hotelhstdf

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// TaobaoXhotelHotelMessageReceive 接收道消息接口
// taobao.xhotel.hotel.message.receive
//
// 接收道消息接口
func TaobaoXhotelHotelMessageReceive(clt *core.SDKClient, req *hotelhstdf.TaobaoXhotelHotelMessageReceiveAPIRequest, session string) (*hotelhstdf.TaobaoXhotelHotelMessageReceiveAPIResponse, error) {
	var resp hotelhstdf.TaobaoXhotelHotelMessageReceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
