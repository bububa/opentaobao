package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取海王用户权限包 APIRequest
alibaba.seaking.servicepack

获取海王用户权限包
*/
type AlibabaSeakingServicepackRequest struct {
    model.Params

    // 验证类型
    identifyType   string 

    // 验证类型下的唯一id
    identifier   string 

}

func NewAlibabaSeakingServicepackRequest() *AlibabaSeakingServicepackRequest{
    return &AlibabaSeakingServicepackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSeakingServicepackRequest) GetApiMethodName() string {
    return "alibaba.seaking.servicepack"
}

func (r AlibabaSeakingServicepackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSeakingServicepackRequest) SetIdentifyType(identifyType string) error {
    r.identifyType = identifyType
    r.Set("identify_type", identifyType)
    return nil
}

func (r AlibabaSeakingServicepackRequest) GetIdentifyType() string {
    return r.identifyType
}

func (r *AlibabaSeakingServicepackRequest) SetIdentifier(identifier string) error {
    r.identifier = identifier
    r.Set("identifier", identifier)
    return nil
}

func (r AlibabaSeakingServicepackRequest) GetIdentifier() string {
    return r.identifier
}

