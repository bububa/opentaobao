package ioti

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对混合云提供的刷图接口 APIRequest
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

func NewAlibabaItEslEslimageShowimagecommonRequest() *AlibabaItEslEslimageShowimagecommonRequest{
    return &AlibabaItEslEslimageShowimagecommonRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaItEslEslimageShowimagecommonRequest) GetApiMethodName() string {
    return "alibaba.it.esl.eslimage.showimagecommon"
}

func (r AlibabaItEslEslimageShowimagecommonRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaItEslEslimageShowimagecommonRequest) SetMac(mac string) error {
    r.mac = mac
    r.Set("mac", mac)
    return nil
}

func (r AlibabaItEslEslimageShowimagecommonRequest) GetMac() string {
    return r.mac
}

func (r *AlibabaItEslEslimageShowimagecommonRequest) SetContent2(content2 string) error {
    r.content2 = content2
    r.Set("content2", content2)
    return nil
}

func (r AlibabaItEslEslimageShowimagecommonRequest) GetContent2() string {
    return r.content2
}

func (r *AlibabaItEslEslimageShowimagecommonRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r AlibabaItEslEslimageShowimagecommonRequest) GetContent() string {
    return r.content
}

func (r *AlibabaItEslEslimageShowimagecommonRequest) SetIsCompress(isCompress string) error {
    r.isCompress = isCompress
    r.Set("is_compress", isCompress)
    return nil
}

func (r AlibabaItEslEslimageShowimagecommonRequest) GetIsCompress() string {
    return r.isCompress
}

func (r *AlibabaItEslEslimageShowimagecommonRequest) SetIsManual(isManual string) error {
    r.isManual = isManual
    r.Set("is_manual", isManual)
    return nil
}

func (r AlibabaItEslEslimageShowimagecommonRequest) GetIsManual() string {
    return r.isManual
}

