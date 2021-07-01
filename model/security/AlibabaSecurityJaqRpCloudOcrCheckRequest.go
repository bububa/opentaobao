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
type AlibabaSecurityJaqRpCloudOcrCheckAPIRequest struct {
    model.Params
    // token
    _verifyToken   string
    // 要识别的信息
    _imageUrls   string
}

// 初始化AlibabaSecurityJaqRpCloudOcrCheckAPIRequest对象
func NewAlibabaSecurityJaqRpCloudOcrCheckRequest() *AlibabaSecurityJaqRpCloudOcrCheckAPIRequest{
    return &AlibabaSecurityJaqRpCloudOcrCheckAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.ocr.check"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) SetVerifyToken(_verifyToken string) error {
    r._verifyToken = _verifyToken
    r.Set("verify_token", _verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) GetVerifyToken() string {
    return r._verifyToken
}
// ImageUrls Setter
// 要识别的信息
func (r *AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) SetImageUrls(_imageUrls string) error {
    r._imageUrls = _imageUrls
    r.Set("image_urls", _imageUrls)
    return nil
}

// ImageUrls Getter
func (r AlibabaSecurityJaqRpCloudOcrCheckAPIRequest) GetImageUrls() string {
    return r._imageUrls
}
