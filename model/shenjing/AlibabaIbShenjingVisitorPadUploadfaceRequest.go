package shenjing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
访客PAD上传人脸 API请求
alibaba.ib.shenjing.visitor.pad.uploadface

访客PAD端上传人脸。
*/
type AlibabaIbShenjingVisitorPadUploadfaceRequest struct {
    model.Params
    // 访客ID
    _id   string
    // 图片URL
    _image   string
}

// 初始化AlibabaIbShenjingVisitorPadUploadfaceRequest对象
func NewAlibabaIbShenjingVisitorPadUploadfaceRequest() *AlibabaIbShenjingVisitorPadUploadfaceRequest{
    return &AlibabaIbShenjingVisitorPadUploadfaceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIbShenjingVisitorPadUploadfaceRequest) GetApiMethodName() string {
    return "alibaba.ib.shenjing.visitor.pad.uploadface"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIbShenjingVisitorPadUploadfaceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 访客ID
func (r *AlibabaIbShenjingVisitorPadUploadfaceRequest) SetId(_id string) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaIbShenjingVisitorPadUploadfaceRequest) GetId() string {
    return r._id
}
// Image Setter
// 图片URL
func (r *AlibabaIbShenjingVisitorPadUploadfaceRequest) SetImage(_image string) error {
    r._image = _image
    r.Set("image", _image)
    return nil
}

// Image Getter
func (r AlibabaIbShenjingVisitorPadUploadfaceRequest) GetImage() string {
    return r._image
}
