package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScShopConvert 淘宝客-服务商-店铺链接转换
// taobao.tbk.sc.shop.convert
//
// 服务商使用。支持入参推广者对应的“推广位”和卖家id，获取对应的店铺推广链接。
func TaobaoTbkScShopConvert(clt *core.SDKClient, req *tbk.TaobaoTbkScShopConvertAPIRequest, session string) (*tbk.TaobaoTbkScShopConvertAPIResponse, error) {
	var resp tbk.TaobaoTbkScShopConvertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
