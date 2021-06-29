package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ocr同时实名校验 API请求
alibaba.security.jaq.rp.ocr.check

聚安全实人认证证件OCR识别功能接口
*/
type AlibabaSecurityJaqRpOcrCheckRequest struct {
    model.Params
    // token
    _verifyToken   string
    // 要识别的信息
    _imageUrls   string
}

// 初始化AlibabaSecurityJaqRpOcrCheckRequest对象
func NewAlibabaSecurityJaqRpOcrCheckRequest() *AlibabaSecurityJaqRpOcrCheckRequest{
    return &AlibabaSecurityJaqRpOcrCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpOcrCheckRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.ocr.check"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpOcrCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpOcrCheckRequest) SetVerifyToken(_verifyToken string) error {
    r._verifyToken = _verifyToken
    r.Set("verify_token", _verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpOcrCheckRequest) GetVerifyToken() string {
    return r._verifyToken
}
// ImageUrls Setter
// 要识别的信息
func (r *AlibabaSecurityJaqRpOcrCheckRequest) SetImageUrls(_imageUrls string) error {
    r._imageUrls = _imageUrls
    r.Set("image_urls", _imageUrls)
    return nil
}

// ImageUrls Getter
func (r AlibabaSecurityJaqRpOcrCheckRequest) GetImageUrls() string {
    return r._imageUrls
}
