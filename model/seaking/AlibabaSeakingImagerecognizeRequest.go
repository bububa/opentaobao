package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片语种识别 API请求
alibaba.seaking.imagerecognize

图片语种识别
*/
type AlibabaSeakingImagerecognizeRequest struct {
    model.Params
    // 扩展信息
    extra   *Extra
    // erp用户id
    identifier   string
    // 调用来源(erp名称)
    identifierType   string
    // 图片url
    url   string
}

// 初始化AlibabaSeakingImagerecognizeRequest对象
func NewAlibabaSeakingImagerecognizeRequest() *AlibabaSeakingImagerecognizeRequest{
    return &AlibabaSeakingImagerecognizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingImagerecognizeRequest) GetApiMethodName() string {
    return "alibaba.seaking.imagerecognize"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingImagerecognizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Extra Setter
// 扩展信息
func (r *AlibabaSeakingImagerecognizeRequest) SetExtra(extra *Extra) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

// Extra Getter
func (r AlibabaSeakingImagerecognizeRequest) GetExtra() *Extra {
    return r.extra
}
// Identifier Setter
// erp用户id
func (r *AlibabaSeakingImagerecognizeRequest) SetIdentifier(identifier string) error {
    r.identifier = identifier
    r.Set("identifier", identifier)
    return nil
}

// Identifier Getter
func (r AlibabaSeakingImagerecognizeRequest) GetIdentifier() string {
    return r.identifier
}
// IdentifierType Setter
// 调用来源(erp名称)
func (r *AlibabaSeakingImagerecognizeRequest) SetIdentifierType(identifierType string) error {
    r.identifierType = identifierType
    r.Set("identifier_type", identifierType)
    return nil
}

// IdentifierType Getter
func (r AlibabaSeakingImagerecognizeRequest) GetIdentifierType() string {
    return r.identifierType
}
// Url Setter
// 图片url
func (r *AlibabaSeakingImagerecognizeRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

// Url Getter
func (r AlibabaSeakingImagerecognizeRequest) GetUrl() string {
    return r.url
}
