package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusdisableqrcodeset 自助机失效二维码
// taobao.bus.disableqrcode.set
//
// 使创建的二维码失效
func Taobaobusdisableqrcodeset(clt *core.SDKClient, req *bus.TaobaobusdisableqrcodesetAPIRequest, session string) (*bus.TaobaobusdisableqrcodesetAPIResponse, error) {
	var resp bus.TaobaobusdisableqrcodesetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
