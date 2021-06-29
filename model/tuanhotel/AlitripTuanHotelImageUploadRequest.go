package tuanhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片上传接口 API请求
alitrip.tuan.hotel.image.upload

用户调用此接口完成外网图片上传至卖家图片中心，此接口返回图片中心的图片地址
*/
type AlitripTuanHotelImageUploadRequest struct {
    model.Params
    // 上传图片信息列表，最多一次支持5张图片上传。单张图片大小限制为1M
    _imageInfoList   []ImageInfoVOList
}

// 初始化AlitripTuanHotelImageUploadRequest对象
func NewAlitripTuanHotelImageUploadRequest() *AlitripTuanHotelImageUploadRequest{
    return &AlitripTuanHotelImageUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTuanHotelImageUploadRequest) GetApiMethodName() string {
    return "alitrip.tuan.hotel.image.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTuanHotelImageUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImageInfoList Setter
// 上传图片信息列表，最多一次支持5张图片上传。单张图片大小限制为1M
func (r *AlitripTuanHotelImageUploadRequest) SetImageInfoList(_imageInfoList []ImageInfoVOList) error {
    r._imageInfoList = _imageInfoList
    r.Set("image_info_list", _imageInfoList)
    return nil
}

// ImageInfoList Getter
func (r AlitripTuanHotelImageUploadRequest) GetImageInfoList() []ImageInfoVOList {
    return r._imageInfoList
}
