package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusTvmcreateqrcodeSet 自助机生成支付宝支付二维码
// taobao.bus.tvmcreateqrcode.set
//
// 用于汽车票线下自助机调用获取支付宝的二维码
func TaobaoBusTvmcreateqrcodeSet(clt *core.SDKClient, req *bus.TaobaoBusTvmcreateqrcodeSetAPIRequest, resp *bus.TaobaoBusTvmcreateqrcodeSetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
