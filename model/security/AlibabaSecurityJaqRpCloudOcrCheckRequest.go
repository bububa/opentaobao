package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ocr同时实名校验 API请求
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

// 初始化AlibabaSecurityJaqRpCloudOcrCheckRequest对象
func NewAlibabaSecurityJaqRpCloudOcrCheckRequest() *AlibabaSecurityJaqRpCloudOcrCheckRequest{
    return &AlibabaSecurityJaqRpCloudOcrCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudOcrCheckRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.ocr.check"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudOcrCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpCloudOcrCheckRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudOcrCheckRequest) GetVerifyToken() string {
    return r.verifyToken
}
// ImageUrls Setter
// 要识别的信息
func (r *AlibabaSecurityJaqRpCloudOcrCheckRequest) SetImageUrls(imageUrls string) error {
    r.imageUrls = imageUrls
    r.Set("image_urls", imageUrls)
    return nil
}

// ImageUrls Getter
func (r AlibabaSecurityJaqRpCloudOcrCheckRequest) GetImageUrls() string {
    return r.imageUrls
}
