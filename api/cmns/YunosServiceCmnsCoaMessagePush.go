package cmns

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// Yunosservicecmnscoamessagepush 消息推送接口
// yunos.service.cmns.coa.message.push
//
// 调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
func Yunosservicecmnscoamessagepush(clt *core.SDKClient, req *cmns.YunosservicecmnscoamessagepushAPIRequest, session string) (*cmns.YunosservicecmnscoamessagepushAPIResponse, error) {
	var resp cmns.YunosservicecmnscoamessagepushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
