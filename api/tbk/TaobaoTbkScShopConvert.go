package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkscshopconvert 淘宝客-服务商-店铺链接转换
// taobao.tbk.sc.shop.convert
//
// 服务商使用。支持入参推广者对应的“推广位”和卖家id，获取对应的店铺推广链接。
func Taobaotbkscshopconvert(clt *core.SDKClient, req *tbk.TaobaotbkscshopconvertAPIRequest, session string) (*tbk.TaobaotbkscshopconvertAPIResponse, error) {
	var resp tbk.TaobaotbkscshopconvertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
