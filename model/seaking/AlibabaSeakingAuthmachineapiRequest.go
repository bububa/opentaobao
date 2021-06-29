package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
机翻Api授权 APIRequest
alibaba.seaking.authmachineapi

机翻Api授权
*/
type AlibabaSeakingAuthmachineapiRequest struct {
    model.Params

    // erp名称
    identifyType   string 

    // erp用户id
    identifier   string 

    // 店铺所在平台
    subIdentifyType   string 

    // 店铺id(ae为cn开头的店铺id, lazada为邮箱)
    subIdentifier   string 

}

func NewAlibabaSeakingAuthmachineapiRequest() *AlibabaSeakingAuthmachineapiRequest{
    return &AlibabaSeakingAuthmachineapiRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSeakingAuthmachineapiRequest) GetApiMethodName() string {
    return "alibaba.seaking.authmachineapi"
}

func (r AlibabaSeakingAuthmachineapiRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSeakingAuthmachineapiRequest) SetIdentifyType(identifyType string) error {
    r.identifyType = identifyType
    r.Set("identify_type", identifyType)
    return nil
}

func (r AlibabaSeakingAuthmachineapiRequest) GetIdentifyType() string {
    return r.identifyType
}

func (r *AlibabaSeakingAuthmachineapiRequest) SetIdentifier(identifier string) error {
    r.identifier = identifier
    r.Set("identifier", identifier)
    return nil
}

func (r AlibabaSeakingAuthmachineapiRequest) GetIdentifier() string {
    return r.identifier
}

func (r *AlibabaSeakingAuthmachineapiRequest) SetSubIdentifyType(subIdentifyType string) error {
    r.subIdentifyType = subIdentifyType
    r.Set("sub_identify_type", subIdentifyType)
    return nil
}

func (r AlibabaSeakingAuthmachineapiRequest) GetSubIdentifyType() string {
    return r.subIdentifyType
}

func (r *AlibabaSeakingAuthmachineapiRequest) SetSubIdentifier(subIdentifier string) error {
    r.subIdentifier = subIdentifier
    r.Set("sub_identifier", subIdentifier)
    return nil
}

func (r AlibabaSeakingAuthmachineapiRequest) GetSubIdentifier() string {
    return r.subIdentifier
}

