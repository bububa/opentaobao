package cmns

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// YunosServiceCmnsCoaPush 消息推送接口
// yunos.service.cmns.coa.push
//
// 调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
func YunosServiceCmnsCoaPush(clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaPushAPIRequest, resp *cmns.YunosServiceCmnsCoaPushAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
