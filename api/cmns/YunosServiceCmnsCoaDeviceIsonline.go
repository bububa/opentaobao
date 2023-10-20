package cmns

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// YunosServiceCmnsCoaDeviceIsonline 根据设备id查询设备是否在线
// yunos.service.cmns.coa.device.isonline
//
// 根据设备id查询设备是否在线
func YunosServiceCmnsCoaDeviceIsonline(clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaDeviceIsonlineAPIRequest, resp *cmns.YunosServiceCmnsCoaDeviceIsonlineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
