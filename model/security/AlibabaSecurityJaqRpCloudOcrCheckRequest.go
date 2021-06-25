package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ocr同时实名校验 APIRequest
alibaba.security.jaq.rp.cloud.ocr.check

聚安全实人认证证件OCR识别功能接口
*/
type AlibabaSecurityJaqRpCloudOcrCheckRequest struct {
    model.Params

    // token
    verifyToken   string 

    // 要识别的信息
    imageUrls   string 

}

func NewAlibabaSecurityJaqRpCloudOcrCheckRequest() *AlibabaSecurityJaqRpCloudOcrCheckRequest{
    return &AlibabaSecurityJaqRpCloudOcrCheckRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpCloudOcrCheckRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.ocr.check"
}

func (r AlibabaSecurityJaqRpCloudOcrCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpCloudOcrCheckRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

func (r AlibabaSecurityJaqRpCloudOcrCheckRequest) GetVerifyToken() string {
    return r.verifyToken
}

func (r *AlibabaSecurityJaqRpCloudOcrCheckRequest) SetImageUrls(imageUrls string) error {
    r.imageUrls = imageUrls
    r.Set("image_urls", imageUrls)
    return nil
}

func (r AlibabaSecurityJaqRpCloudOcrCheckRequest) GetImageUrls() string {
    return r.imageUrls
}

