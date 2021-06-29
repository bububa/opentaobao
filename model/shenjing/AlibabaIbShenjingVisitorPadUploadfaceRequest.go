package shenjing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
访客PAD上传人脸 APIRequest
alibaba.ib.shenjing.visitor.pad.uploadface

访客PAD端上传人脸。
*/
type AlibabaIbShenjingVisitorPadUploadfaceRequest struct {
    model.Params

    // 访客ID
    id   string 

    // 图片URL
    image   string 

}

func NewAlibabaIbShenjingVisitorPadUploadfaceRequest() *AlibabaIbShenjingVisitorPadUploadfaceRequest{
    return &AlibabaIbShenjingVisitorPadUploadfaceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIbShenjingVisitorPadUploadfaceRequest) GetApiMethodName() string {
    return "alibaba.ib.shenjing.visitor.pad.uploadface"
}

func (r AlibabaIbShenjingVisitorPadUploadfaceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIbShenjingVisitorPadUploadfaceRequest) SetId(id string) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaIbShenjingVisitorPadUploadfaceRequest) GetId() string {
    return r.id
}

func (r *AlibabaIbShenjingVisitorPadUploadfaceRequest) SetImage(image string) error {
    r.image = image
    r.Set("image", image)
    return nil
}

func (r AlibabaIbShenjingVisitorPadUploadfaceRequest) GetImage() string {
    return r.image
}

