package cmns

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cmns"
)

// YunosServiceCmnsCoaDeviceGet 设备详情查询
// yunos.service.cmns.coa.device.get
//
// 第三方应用开发者调用此接口查询设备详情
func YunosServiceCmnsCoaDeviceGet(clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaDeviceGetAPIRequest, resp *cmns.YunosServiceCmnsCoaDeviceGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
