package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

/* TaobaoWirelessBuntingShopShorturlCreate
通过店铺id取得短链
taobao.wireless.bunting.shop.shorturl.create

通过店铺id取得短链 */
func TaobaoWirelessBuntingShopShorturlCreate(clt *core.SDKClient, req *mtopopen.TaobaoWirelessBuntingShopShorturlCreateAPIRequest, session string) (*mtopopen.TaobaoWirelessBuntingShopShorturlCreateAPIResponse, error) {
	var resp mtopopen.TaobaoWirelessBuntingShopShorturlCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
