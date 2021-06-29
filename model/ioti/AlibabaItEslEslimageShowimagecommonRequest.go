package ioti

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对混合云提供的刷图接口 API请求
alibaba.it.esl.eslimage.showimagecommon

混合云使用，提供给isv和我们混合云环境部署的应用刷图
*/
type AlibabaItEslEslimageShowimagecommonRequest struct {
    model.Params
    // ma地址
    _mac   string
    // 图片的base64编码,图片要和价签大小一致
    _content2   string
    // 图片2的base64编码,图片要和价签大小一致
    _content   string
    // 是否压缩，默认不传，字符串：yes，no
    _isCompress   string
    // 是否手动刷图，默认不传，字符串：true，false
    _isManual   string
}

// 初始化AlibabaItEslEslimageShowimagecommonRequest对象
func NewAlibabaItEslEslimageShowimagecommonRequest() *AlibabaItEslEslimageShowimagecommonRequest{
    return &AlibabaItEslEslimageShowimagecommonRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItEslEslimageShowimagecommonRequest) GetApiMethodName() string {
    return "alibaba.it.esl.eslimage.showimagecommon"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItEslEslimageShowimagecommonRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mac Setter
// ma地址
func (r *AlibabaItEslEslimageShowimagecommonRequest) SetMac(_mac string) error {
    r._mac = _mac
    r.Set("mac", _mac)
    return nil
}

// Mac Getter
func (r AlibabaItEslEslimageShowimagecommonRequest) GetMac() string {
    return r._mac
}
// Content2 Setter
// 图片的base64编码,图片要和价签大小一致
func (r *AlibabaItEslEslimageShowimagecommonRequest) SetContent2(_content2 string) error {
    r._content2 = _content2
    r.Set("content2", _content2)
    return nil
}

// Content2 Getter
func (r AlibabaItEslEslimageShowimagecommonRequest) GetContent2() string {
    return r._content2
}
// Content Setter
// 图片2的base64编码,图片要和价签大小一致
func (r *AlibabaItEslEslimageShowimagecommonRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r AlibabaItEslEslimageShowimagecommonRequest) GetContent() string {
    return r._content
}
// IsCompress Setter
// 是否压缩，默认不传，字符串：yes，no
func (r *AlibabaItEslEslimageShowimagecommonRequest) SetIsCompress(_isCompress string) error {
    r._isCompress = _isCompress
    r.Set("is_compress", _isCompress)
    return nil
}

// IsCompress Getter
func (r AlibabaItEslEslimageShowimagecommonRequest) GetIsCompress() string {
    return r._isCompress
}
// IsManual Setter
// 是否手动刷图，默认不传，字符串：true，false
func (r *AlibabaItEslEslimageShowimagecommonRequest) SetIsManual(_isManual string) error {
    r._isManual = _isManual
    r.Set("is_manual", _isManual)
    return nil
}

// IsManual Getter
func (r AlibabaItEslEslimageShowimagecommonRequest) GetIsManual() string {
    return r._isManual
}
