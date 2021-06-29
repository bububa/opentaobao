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
    mac   string
    // 图片的base64编码,图片要和价签大小一致
    content2   string
    // 图片2的base64编码,图片要和价签大小一致
    content   string
    // 是否压缩，默认不传，字符串：yes，no
    isCompress   string
    // 是否手动刷图，默认不传，字符串：true，false
    isManual   string
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
func (r *AlibabaItEslEslimageShowimagecommonRequest) SetMac(mac string) error {
    r.mac = mac
    r.Set("mac", mac)
    return nil
}

// Mac Getter
func (r AlibabaItEslEslimageShowimagecommonRequest) GetMac() string {
    return r.mac
}
// Content2 Setter
// 图片的base64编码,图片要和价签大小一致
func (r *AlibabaItEslEslimageShowimagecommonRequest) SetContent2(content2 string) error {
    r.content2 = content2
    r.Set("content2", content2)
    return nil
}

// Content2 Getter
func (r AlibabaItEslEslimageShowimagecommonRequest) GetContent2() string {
    return r.content2
}
// Content Setter
// 图片2的base64编码,图片要和价签大小一致
func (r *AlibabaItEslEslimageShowimagecommonRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r AlibabaItEslEslimageShowimagecommonRequest) GetContent() string {
    return r.content
}
// IsCompress Setter
// 是否压缩，默认不传，字符串：yes，no
func (r *AlibabaItEslEslimageShowimagecommonRequest) SetIsCompress(isCompress string) error {
    r.isCompress = isCompress
    r.Set("is_compress", isCompress)
    return nil
}

// IsCompress Getter
func (r AlibabaItEslEslimageShowimagecommonRequest) GetIsCompress() string {
    return r.isCompress
}
// IsManual Setter
// 是否手动刷图，默认不传，字符串：true，false
func (r *AlibabaItEslEslimageShowimagecommonRequest) SetIsManual(isManual string) error {
    r.isManual = isManual
    r.Set("is_manual", isManual)
    return nil
}

// IsManual Getter
func (r AlibabaItEslEslimageShowimagecommonRequest) GetIsManual() string {
    return r.isManual
}
