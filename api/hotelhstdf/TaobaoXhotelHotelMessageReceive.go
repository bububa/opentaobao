package hotelhstdf

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// Taobaoxhotelhotelmessagereceive 接收道消息接口
// taobao.xhotel.hotel.message.receive
//
// 接收道消息接口
func Taobaoxhotelhotelmessagereceive(clt *core.SDKClient, req *hotelhstdf.TaobaoxhotelhotelmessagereceiveAPIRequest, session string) (*hotelhstdf.TaobaoxhotelhotelmessagereceiveAPIResponse, error) {
	var resp hotelhstdf.TaobaoxhotelhotelmessagereceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
