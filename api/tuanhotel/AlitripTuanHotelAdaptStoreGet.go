package tuanhotel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// AlitripTuanHotelAdaptStoreGet 酒店团购套餐关联适用门店
// alitrip.tuan.hotel.adapt.store.get
//
// 输入shid，返回关联门店详情信息
func AlitripTuanHotelAdaptStoreGet(ctx context.Context, clt *core.SDKClient, req *tuanhotel.AlitripTuanHotelAdaptStoreGetAPIRequest, resp *tuanhotel.AlitripTuanHotelAdaptStoreGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
