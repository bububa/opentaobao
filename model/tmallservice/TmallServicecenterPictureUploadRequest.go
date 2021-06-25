package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传图片 APIRequest
tmall.servicecenter.picture.upload

给服务商ERP系统使用，用于上传图片保存在天猫，一般用于工单信息回传时候保存服务商的服务证明信息相关的图片。
*/
type TmallServicecenterPictureUploadRequest struct {
    model.Params

    // 图片文件二进制流
    img   []byte 

    // 图片全称包括扩展名。目前支持 jpg jpeg png
    pictureName   string 

    // true返回Https地址
    isHttps   bool 

}

func NewTmallServicecenterPictureUploadRequest() *TmallServicecenterPictureUploadRequest{
    return &TmallServicecenterPictureUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterPictureUploadRequest) GetApiMethodName() string {
    return "tmall.servicecenter.picture.upload"
}

func (r TmallServicecenterPictureUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterPictureUploadRequest) SetImg(img []byte) error {
    r.img = img
    r.Set("img", img)
    return nil
}

func (r TmallServicecenterPictureUploadRequest) GetImg() []byte {
    return r.img
}

func (r *TmallServicecenterPictureUploadRequest) SetPictureName(pictureName string) error {
    r.pictureName = pictureName
    r.Set("picture_name", pictureName)
    return nil
}

func (r TmallServicecenterPictureUploadRequest) GetPictureName() string {
    return r.pictureName
}

func (r *TmallServicecenterPictureUploadRequest) SetIsHttps(isHttps bool) error {
    r.isHttps = isHttps
    r.Set("is_https", isHttps)
    return nil
}

func (r TmallServicecenterPictureUploadRequest) GetIsHttps() bool {
    return r.isHttps
}

