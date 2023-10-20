package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// AlitripTuanHotelItemSkuUpdate 酒店团购套餐商品SKU更新和新增
// alitrip.tuan.hotel.item.sku.update
//
// 商户对发布的宝贝套餐价格库存信息进行更新，对于已存在的sku，未进行传递则不会进行覆盖。skuId必须为已存在的skuId，暂不支持库存类型的更改。因发布页改造升级，2020.03.05将下线此接口的新增SKU功能，更新SKU功能将保留，但商户2020.03.05后须前往发布页进行宝贝更新后，方可调用本接口。对于日历库存宝贝日历维度的价格和库存数据的更新，此接口存在调用超时的问题，不推荐使用，若有诉求，请使用alitrip.tuan.hotel.item.sku.calendar.update接口（该接口提供增量更新能力），接口地址为https://open.taobao.com/api.htm?docId=48160&amp;docType=2&amp;scopeId=12326
func AlitripTuanHotelItemSkuUpdate(clt *core.SDKClient, req *tuanhotel.AlitripTuanHotelItemSkuUpdateAPIRequest, session string) (*tuanhotel.AlitripTuanHotelItemSkuUpdateAPIResponse, error) {
	var resp tuanhotel.AlitripTuanHotelItemSkuUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
