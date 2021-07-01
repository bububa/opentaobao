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
type AlibabaSeakingImagerecognizeAPIRequest struct {
    model.Params
    // 扩展信息
    _extra   *Extra
    // erp用户id
    _identifier   string
    // 调用来源(erp名称)
    _identifierType   string
    // 图片url
    _url   string
}

// 初始化AlibabaSeakingImagerecognizeAPIRequest对象
func NewAlibabaSeakingImagerecognizeRequest() *AlibabaSeakingImagerecognizeAPIRequest{
    return &AlibabaSeakingImagerecognizeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingImagerecognizeAPIRequest) GetApiMethodName() string {
    return "alibaba.seaking.imagerecognize"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingImagerecognizeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Extra Setter
// 扩展信息
func (r *AlibabaSeakingImagerecognizeAPIRequest) SetExtra(_extra *Extra) error {
    r._extra = _extra
    r.Set("extra", _extra)
    return nil
}

// Extra Getter
func (r AlibabaSeakingImagerecognizeAPIRequest) GetExtra() *Extra {
    return r._extra
}
// Identifier Setter
// erp用户id
func (r *AlibabaSeakingImagerecognizeAPIRequest) SetIdentifier(_identifier string) error {
    r._identifier = _identifier
    r.Set("identifier", _identifier)
    return nil
}

// Identifier Getter
func (r AlibabaSeakingImagerecognizeAPIRequest) GetIdentifier() string {
    return r._identifier
}
// IdentifierType Setter
// 调用来源(erp名称)
func (r *AlibabaSeakingImagerecognizeAPIRequest) SetIdentifierType(_identifierType string) error {
    r._identifierType = _identifierType
    r.Set("identifier_type", _identifierType)
    return nil
}

// IdentifierType Getter
func (r AlibabaSeakingImagerecognizeAPIRequest) GetIdentifierType() string {
    return r._identifierType
}
// Url Setter
// 图片url
func (r *AlibabaSeakingImagerecognizeAPIRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r AlibabaSeakingImagerecognizeAPIRequest) GetUrl() string {
    return r._url
}
