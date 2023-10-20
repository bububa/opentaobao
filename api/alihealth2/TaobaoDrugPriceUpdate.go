package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Taobaodrugpriceupdate 商家更新宝贝价格
// taobao.drug.price.update
//
// 商家更新价格
func Taobaodrugpriceupdate(clt *core.SDKClient, req *alihealth2.TaobaodrugpriceupdateAPIRequest, session string) (*alihealth2.TaobaodrugpriceupdateAPIResponse, error) {
	var resp alihealth2.TaobaodrugpriceupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
