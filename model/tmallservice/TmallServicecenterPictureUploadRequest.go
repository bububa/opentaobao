package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传图片 API请求
tmall.servicecenter.picture.upload

给服务商ERP系统使用，用于上传图片保存在天猫，一般用于工单信息回传时候保存服务商的服务证明信息相关的图片。
*/
type TmallServicecenterPictureUploadRequest struct {
    model.Params
    // 图片文件二进制流
    _img   []*model.File
    // 图片全称包括扩展名。目前支持 jpg jpeg png
    _pictureName   string
    // true返回Https地址
    _isHttps   bool
}

// 初始化TmallServicecenterPictureUploadRequest对象
func NewTmallServicecenterPictureUploadRequest() *TmallServicecenterPictureUploadRequest{
    return &TmallServicecenterPictureUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterPictureUploadRequest) GetApiMethodName() string {
    return "tmall.servicecenter.picture.upload"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterPictureUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Img Setter
// 图片文件二进制流
func (r *TmallServicecenterPictureUploadRequest) SetImg(_img []*model.File) error {
    r._img = _img
    r.Set("img", _img)
    return nil
}

// Img Getter
func (r TmallServicecenterPictureUploadRequest) GetImg() []*model.File {
    return r._img
}
// PictureName Setter
// 图片全称包括扩展名。目前支持 jpg jpeg png
func (r *TmallServicecenterPictureUploadRequest) SetPictureName(_pictureName string) error {
    r._pictureName = _pictureName
    r.Set("picture_name", _pictureName)
    return nil
}

// PictureName Getter
func (r TmallServicecenterPictureUploadRequest) GetPictureName() string {
    return r._pictureName
}
// IsHttps Setter
// true返回Https地址
func (r *TmallServicecenterPictureUploadRequest) SetIsHttps(_isHttps bool) error {
    r._isHttps = _isHttps
    r.Set("is_https", _isHttps)
    return nil
}

// IsHttps Getter
func (r TmallServicecenterPictureUploadRequest) GetIsHttps() bool {
    return r._isHttps
}
