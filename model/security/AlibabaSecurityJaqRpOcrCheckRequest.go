package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ocr同时实名校验 APIRequest
alibaba.security.jaq.rp.ocr.check

聚安全实人认证证件OCR识别功能接口
*/
type AlibabaSecurityJaqRpOcrCheckRequest struct {
    model.Params

    // token
    verifyToken   string 

    // 要识别的信息
    imageUrls   string 

}

func NewAlibabaSecurityJaqRpOcrCheckRequest() *AlibabaSecurityJaqRpOcrCheckRequest{
    return &AlibabaSecurityJaqRpOcrCheckRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpOcrCheckRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.ocr.check"
}

func (r AlibabaSecurityJaqRpOcrCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpOcrCheckRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

func (r AlibabaSecurityJaqRpOcrCheckRequest) GetVerifyToken() string {
    return r.verifyToken
}

func (r *AlibabaSecurityJaqRpOcrCheckRequest) SetImageUrls(imageUrls string) error {
    r.imageUrls = imageUrls
    r.Set("image_urls", imageUrls)
    return nil
}

func (r AlibabaSecurityJaqRpOcrCheckRequest) GetImageUrls() string {
    return r.imageUrls
}

