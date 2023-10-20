package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautoorderqrcode 根据商品id列表获取可扫描下单二维码
// tmall.aliauto.order.qrcode
//
// 根据商品id列表获取可扫描下单二维码
func Tmallaliautoorderqrcode(clt *core.SDKClient, req *tmallcar.TmallaliautoorderqrcodeAPIRequest, session string) (*tmallcar.TmallaliautoorderqrcodeAPIResponse, error) {
	var resp tmallcar.TmallaliautoorderqrcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
