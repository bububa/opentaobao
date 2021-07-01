package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTuanHotelImageUploadAPIRequest
图片上传接口 API请求
alitrip.tuan.hotel.image.upload

用户调用此接口完成外网图片上传至卖家图片中心，此接口返回图片中心的图片地址 */
type AlitripTuanHotelImageUploadAPIRequest struct {
	model.Params
	// 上传图片信息列表，最多一次支持5张图片上传。单张图片大小限制为1M
	_imageInfoList []ImageInfoVOList
}

// New
