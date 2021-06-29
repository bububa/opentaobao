package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证证件OCR识别功能接口 API请求
alibaba.security.jaq.rp.ocr

聚安全实人认证证件OCR识别功能接口
*/
type AlibabaSecurityJaqRpOcrRequest struct {
    model.Params
    // token
    _verifyToken   string
    // 要识别的信息
    _imageUrls   string
}

// 初始化AlibabaSecurityJaqRpOcrRequest对象
func NewAlibabaSecurityJaqRpOcrRequest() *AlibabaSecurityJaqRpOcrRequest{
    return &AlibabaSecurityJaqRpOcrRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpOcrRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.ocr"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpOcrRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpOcrRequest) SetVerifyToken(_verifyToken string) error {
    r._verifyToken = _verifyToken
    r.Set("verify_token", _verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpOcrRequest) GetVerifyToken() string {
    return r._verifyToken
}
// ImageUrls Setter
// 要识别的信息
func (r *AlibabaSecurityJaqRpOcrRequest) SetImageUrls(_imageUrls string) error {
    r._imageUrls = _imageUrls
    r.Set("image_urls", _imageUrls)
    return nil
}

// ImageUrls Getter
func (r AlibabaSecurityJaqRpOcrRequest) GetImageUrls() string {
    return r._imageUrls
}
