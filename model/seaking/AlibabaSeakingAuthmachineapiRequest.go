package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
机翻Api授权 API请求
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

// 初始化AlibabaSeakingAuthmachineapiRequest对象
func NewAlibabaSeakingAuthmachineapiRequest() *AlibabaSeakingAuthmachineapiRequest{
    return &AlibabaSeakingAuthmachineapiRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingAuthmachineapiRequest) GetApiMethodName() string {
    return "alibaba.seaking.authmachineapi"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingAuthmachineapiRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IdentifyType Setter
// erp名称
func (r *AlibabaSeakingAuthmachineapiRequest) SetIdentifyType(identifyType string) error {
    r.identifyType = identifyType
    r.Set("identify_type", identifyType)
    return nil
}

// IdentifyType Getter
func (r AlibabaSeakingAuthmachineapiRequest) GetIdentifyType() string {
    return r.identifyType
}
// Identifier Setter
// erp用户id
func (r *AlibabaSeakingAuthmachineapiRequest) SetIdentifier(identifier string) error {
    r.identifier = identifier
    r.Set("identifier", identifier)
    return nil
}

// Identifier Getter
func (r AlibabaSeakingAuthmachineapiRequest) GetIdentifier() string {
    return r.identifier
}
// SubIdentifyType Setter
// 店铺所在平台
func (r *AlibabaSeakingAuthmachineapiRequest) SetSubIdentifyType(subIdentifyType string) error {
    r.subIdentifyType = subIdentifyType
    r.Set("sub_identify_type", subIdentifyType)
    return nil
}

// SubIdentifyType Getter
func (r AlibabaSeakingAuthmachineapiRequest) GetSubIdentifyType() string {
    return r.subIdentifyType
}
// SubIdentifier Setter
// 店铺id(ae为cn开头的店铺id, lazada为邮箱)
func (r *AlibabaSeakingAuthmachineapiRequest) SetSubIdentifier(subIdentifier string) error {
    r.subIdentifier = subIdentifier
    r.Set("sub_identifier", subIdentifier)
    return nil
}

// SubIdentifier Getter
func (r AlibabaSeakingAuthmachineapiRequest) GetSubIdentifier() string {
    return r.subIdentifier
}
