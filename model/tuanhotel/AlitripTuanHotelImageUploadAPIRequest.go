package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelImageUploadAPIRequest 图片上传接口 API请求
// alitrip.tuan.hotel.image.upload
//
// 用户调用此接口完成外网图片上传至卖家图片中心，此接口返回图片中心的图片地址
type AlitripTuanHotelImageUploadAPIRequest struct {
	model.Params
	// 上传图片信息列表，最多一次支持5张图片上传。单张图片大小限制为1M
	_imageInfoList []ImageInfoVolist
}

// NewAlitripTuanHotelImageUploadRequest 初始化AlitripTuanHotelImageUploadAPIRequest对象
func NewAlitripTuanHotelImageUploadRequest() *AlitripTuanHotelImageUploadAPIRequest {
	return &AlitripTuanHotelImageUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTuanHotelImageUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.tuan.hotel.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTuanHotelImageUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTuanHotelImageUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageInfoList is ImageInfoList Setter
// 上传图片信息列表，最多一次支持5张图片上传。单张图片大小限制为1M
func (r *AlitripTuanHotelImageUploadAPIRequest) SetImageInfoList(_imageInfoList []ImageInfoVolist) error {
	r._imageInfoList = _imageInfoList
	r.Set("image_info_list", _imageInfoList)
	return nil
}

// GetImageInfoList ImageInfoList Getter
func (r AlitripTuanHotelImageUploadAPIRequest) GetImageInfoList() []ImageInfoVolist {
	return r._imageInfoList
}
