package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusDisableqrcodeSet 自助机失效二维码
// taobao.bus.disableqrcode.set
//
// 使创建的二维码失效
func TaobaoBusDisableqrcodeSet(clt *core.SDKClient, req *bus.TaobaoBusDisableqrcodeSetAPIRequest, resp *bus.TaobaoBusDisableqrcodeSetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
