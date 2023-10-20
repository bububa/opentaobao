package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// Alitriptuanhotelitemskudelete 酒店团购套餐商品SKU删除
// alitrip.tuan.hotel.item.sku.delete
//
// 商户对发布的宝贝套餐价格库存信息进行删除
func Alitriptuanhotelitemskudelete(clt *core.SDKClient, req *tuanhotel.AlitriptuanhotelitemskudeleteAPIRequest, session string) (*tuanhotel.AlitriptuanhotelitemskudeleteAPIResponse, error) {
	var resp tuanhotel.AlitriptuanhotelitemskudeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
