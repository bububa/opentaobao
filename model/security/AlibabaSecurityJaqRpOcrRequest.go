package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证证件OCR识别功能接口 APIRequest
alibaba.security.jaq.rp.ocr

聚安全实人认证证件OCR识别功能接口
*/
type AlibabaSecurityJaqRpOcrRequest struct {
    model.Params

    // token
    verifyToken   string 

    // 要识别的信息
    imageUrls   string 

}

func NewAlibabaSecurityJaqRpOcrRequest() *AlibabaSecurityJaqRpOcrRequest{
    return &AlibabaSecurityJaqRpOcrRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpOcrRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.ocr"
}

func (r AlibabaSecurityJaqRpOcrRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpOcrRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

func (r AlibabaSecurityJaqRpOcrRequest) GetVerifyToken() string {
    return r.verifyToken
}

func (r *AlibabaSecurityJaqRpOcrRequest) SetImageUrls(imageUrls string) error {
    r.imageUrls = imageUrls
    r.Set("image_urls", imageUrls)
    return nil
}

func (r AlibabaSecurityJaqRpOcrRequest) GetImageUrls() string {
    return r.imageUrls
}

