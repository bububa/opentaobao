package tuanhotel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// AlitripTuanHotelImageUpload 图片上传接口
// alitrip.tuan.hotel.image.upload
//
// 用户调用此接口完成外网图片上传至卖家图片中心，此接口返回图片中心的图片地址
func AlitripTuanHotelImageUpload(ctx context.Context, clt *core.SDKClient, req *tuanhotel.AlitripTuanHotelImageUploadAPIRequest, resp *tuanhotel.AlitripTuanHotelImageUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
