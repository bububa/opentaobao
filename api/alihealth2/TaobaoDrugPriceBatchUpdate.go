package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Taobaodrugpricebatchupdate 商家批量更新宝贝价格
// taobao.drug.price.batch.update
//
// 商家批量更新宝贝价格
func Taobaodrugpricebatchupdate(clt *core.SDKClient, req *alihealth2.TaobaodrugpricebatchupdateAPIRequest, session string) (*alihealth2.TaobaodrugpricebatchupdateAPIResponse, error) {
	var resp alihealth2.TaobaodrugpricebatchupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
