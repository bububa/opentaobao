package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片语种识别 APIRequest
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

func NewAlibabaSeakingImagerecognizeRequest() *AlibabaSeakingImagerecognizeRequest{
    return &AlibabaSeakingImagerecognizeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSeakingImagerecognizeRequest) GetApiMethodName() string {
    return "alibaba.seaking.imagerecognize"
}

func (r AlibabaSeakingImagerecognizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSeakingImagerecognizeRequest) SetExtra(extra *Extra) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

func (r AlibabaSeakingImagerecognizeRequest) GetExtra() *Extra {
    return r.extra
}

func (r *AlibabaSeakingImagerecognizeRequest) SetIdentifier(identifier string) error {
    r.identifier = identifier
    r.Set("identifier", identifier)
    return nil
}

func (r AlibabaSeakingImagerecognizeRequest) GetIdentifier() string {
    return r.identifier
}

func (r *AlibabaSeakingImagerecognizeRequest) SetIdentifierType(identifierType string) error {
    r.identifierType = identifierType
    r.Set("identifier_type", identifierType)
    return nil
}

func (r AlibabaSeakingImagerecognizeRequest) GetIdentifierType() string {
    return r.identifierType
}

func (r *AlibabaSeakingImagerecognizeRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r AlibabaSeakingImagerecognizeRequest) GetUrl() string {
    return r.url
}

