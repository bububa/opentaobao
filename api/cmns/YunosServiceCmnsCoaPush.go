package cmns

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// Yunosservicecmnscoapush 消息推送接口
// yunos.service.cmns.coa.push
//
// 调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
func Yunosservicecmnscoapush(clt *core.SDKClient, req *cmns.YunosservicecmnscoapushAPIRequest, session string) (*cmns.YunosservicecmnscoapushAPIResponse, error) {
	var resp cmns.YunosservicecmnscoapushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
