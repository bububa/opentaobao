package tuanhotel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// AlitripTuanHotelItemInfoGet 宝贝信息查询接口
// alitrip.tuan.hotel.item.info.get
//
// 商家查询发布的宝贝详情信息
func AlitripTuanHotelItemInfoGet(ctx context.Context, clt *core.SDKClient, req *tuanhotel.AlitripTuanHotelItemInfoGetAPIRequest, resp *tuanhotel.AlitripTuanHotelItemInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
