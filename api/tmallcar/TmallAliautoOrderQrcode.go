package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoOrderQrcode 根据商品id列表获取可扫描下单二维码
// tmall.aliauto.order.qrcode
//
// 根据商品id列表获取可扫描下单二维码
func TmallAliautoOrderQrcode(clt *core.SDKClient, req *tmallcar.TmallAliautoOrderQrcodeAPIRequest, resp *tmallcar.TmallAliautoOrderQrcodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
