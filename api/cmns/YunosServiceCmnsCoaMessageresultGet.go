package cmns

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// Yunosservicecmnscoamessageresultget CMNS消息发送到达查询
// yunos.service.cmns.coa.messageresult.get
//
// CMNS消息发送到达查询,根据消息ID查询，仅能查询该appKey所发送的消息
func Yunosservicecmnscoamessageresultget(clt *core.SDKClient, req *cmns.YunosservicecmnscoamessageresultgetAPIRequest, session string) (*cmns.YunosservicecmnscoamessageresultgetAPIResponse, error) {
	var resp cmns.YunosservicecmnscoamessageresultgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
