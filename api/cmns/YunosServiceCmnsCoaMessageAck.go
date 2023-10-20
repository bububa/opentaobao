package cmns

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// Yunosservicecmnscoamessageack 消息回执查询
// yunos.service.cmns.coa.message.ack
//
// 第三方应用开发者调用此接口查询设备是否收到消息，只能查询此appKey床发的消息
func Yunosservicecmnscoamessageack(clt *core.SDKClient, req *cmns.YunosservicecmnscoamessageackAPIRequest, session string) (*cmns.YunosservicecmnscoamessageackAPIResponse, error) {
	var resp cmns.YunosservicecmnscoamessageackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
