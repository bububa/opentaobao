package tuanhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片上传接口 APIRequest
alitrip.tuan.hotel.image.upload

用户调用此接口完成外网图片上传至卖家图片中心，此接口返回图片中心的图片地址
*/
type AlitripTuanHotelImageUploadRequest struct {
    model.Params

    // 上传图片信息列表，最多一次支持5张图片上传。单张图片大小限制为1M
    imageInfoList   []ImageInfoVOList 

}

func NewAlitripTuanHotelImageUploadRequest() *AlitripTuanHotelImageUploadRequest{
    return &AlitripTuanHotelImageUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTuanHotelImageUploadRequest) GetApiMethodName() string {
    return "alitrip.tuan.hotel.image.upload"
}

func (r AlitripTuanHotelImageUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTuanHotelImageUploadRequest) SetImageInfoList(imageInfoList []ImageInfoVOList) error {
    r.imageInfoList = imageInfoList
    r.Set("image_info_list", imageInfoList)
    return nil
}

func (r AlitripTuanHotelImageUploadRequest) GetImageInfoList() []ImageInfoVOList {
    return r.imageInfoList
}

