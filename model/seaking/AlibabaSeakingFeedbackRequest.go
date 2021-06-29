package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
API服务发布成功商品ID回传 APIRequest
alibaba.seaking.feedback

API服务发布成功商品ID回传，用于跟进商品id后续的使用情况
*/
type AlibabaSeakingFeedbackRequest struct {
    model.Params

    // api 接口名字
    invokeApiName   string 

    // 商品投放平台
    platform   string 

    // 商品id
    productId   string 

    // 店铺id(ae为cn开头的店铺id, lazada为邮箱)
    subIdentifier   string 

    // 店铺所在平台
    subIdentifierType   string 

    // erp名称
    identifier   string 

    // erp用户id
    identifierType   string 

}

func NewAlibabaSeakingFeedbackRequest() *AlibabaSeakingFeedbackRequest{
    return &AlibabaSeakingFeedbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSeakingFeedbackRequest) GetApiMethodName() string {
    return "alibaba.seaking.feedback"
}

func (r AlibabaSeakingFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSeakingFeedbackRequest) SetInvokeApiName(invokeApiName string) error {
    r.invokeApiName = invokeApiName
    r.Set("invoke_api_name", invokeApiName)
    return nil
}

func (r AlibabaSeakingFeedbackRequest) GetInvokeApiName() string {
    return r.invokeApiName
}

func (r *AlibabaSeakingFeedbackRequest) SetPlatform(platform string) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

func (r AlibabaSeakingFeedbackRequest) GetPlatform() string {
    return r.platform
}

func (r *AlibabaSeakingFeedbackRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r AlibabaSeakingFeedbackRequest) GetProductId() string {
    return r.productId
}

func (r *AlibabaSeakingFeedbackRequest) SetSubIdentifier(subIdentifier string) error {
    r.subIdentifier = subIdentifier
    r.Set("sub_identifier", subIdentifier)
    return nil
}

func (r AlibabaSeakingFeedbackRequest) GetSubIdentifier() string {
    return r.subIdentifier
}

func (r *AlibabaSeakingFeedbackRequest) SetSubIdentifierType(subIdentifierType string) error {
    r.subIdentifierType = subIdentifierType
    r.Set("sub_identifier_type", subIdentifierType)
    return nil
}

func (r AlibabaSeakingFeedbackRequest) GetSubIdentifierType() string {
    return r.subIdentifierType
}

func (r *AlibabaSeakingFeedbackRequest) SetIdentifier(identifier string) error {
    r.identifier = identifier
    r.Set("identifier", identifier)
    return nil
}

func (r AlibabaSeakingFeedbackRequest) GetIdentifier() string {
    return r.identifier
}

func (r *AlibabaSeakingFeedbackRequest) SetIdentifierType(identifierType string) error {
    r.identifierType = identifierType
    r.Set("identifier_type", identifierType)
    return nil
}

func (r AlibabaSeakingFeedbackRequest) GetIdentifierType() string {
    return r.identifierType
}

