package cmns

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// Yunosservicecmnscoamessageackslist 消息ack记录查询
// yunos.service.cmns.coa.message.acks.list
//
// 第三方应用开发者调用此接口查询消息ack记录
func Yunosservicecmnscoamessageackslist(clt *core.SDKClient, req *cmns.YunosservicecmnscoamessageackslistAPIRequest, session string) (*cmns.YunosservicecmnscoamessageackslistAPIResponse, error) {
	var resp cmns.YunosservicecmnscoamessageackslistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
