package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
API服务发布成功商品ID回传 API请求
alibaba.seaking.feedback

API服务发布成功商品ID回传，用于跟进商品id后续的使用情况
*/
type AlibabaSeakingFeedbackAPIRequest struct {
    model.Params
    // api 接口名字
    _invokeApiName   string
    // 商品投放平台
    _platform   string
    // 商品id
    _productId   string
    // 店铺id(ae为cn开头的店铺id, lazada为邮箱)
    _subIdentifier   string
    // 店铺所在平台
    _subIdentifierType   string
    // erp名称
    _identifier   string
    // erp用户id
    _identifierType   string
}

// 初始化AlibabaSeakingFeedbackAPIRequest对象
func NewAlibabaSeakingFeedbackRequest() *AlibabaSeakingFeedbackAPIRequest{
    return &AlibabaSeakingFeedbackAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingFeedbackAPIRequest) GetApiMethodName() string {
    return "alibaba.seaking.feedback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingFeedbackAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvokeApiName Setter
// api 接口名字
func (r *AlibabaSeakingFeedbackAPIRequest) SetInvokeApiName(_invokeApiName string) error {
    r._invokeApiName = _invokeApiName
    r.Set("invoke_api_name", _invokeApiName)
    return nil
}

// InvokeApiName Getter
func (r AlibabaSeakingFeedbackAPIRequest) GetInvokeApiName() string {
    return r._invokeApiName
}
// Platform Setter
// 商品投放平台
func (r *AlibabaSeakingFeedbackAPIRequest) SetPlatform(_platform string) error {
    r._platform = _platform
    r.Set("platform", _platform)
    return nil
}

// Platform Getter
func (r AlibabaSeakingFeedbackAPIRequest) GetPlatform() string {
    return r._platform
}
// ProductId Setter
// 商品id
func (r *AlibabaSeakingFeedbackAPIRequest) SetProductId(_productId string) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r AlibabaSeakingFeedbackAPIRequest) GetProductId() string {
    return r._productId
}
// SubIdentifier Setter
// 店铺id(ae为cn开头的店铺id, lazada为邮箱)
func (r *AlibabaSeakingFeedbackAPIRequest) SetSubIdentifier(_subIdentifier string) error {
    r._subIdentifier = _subIdentifier
    r.Set("sub_identifier", _subIdentifier)
    return nil
}

// SubIdentifier Getter
func (r AlibabaSeakingFeedbackAPIRequest) GetSubIdentifier() string {
    return r._subIdentifier
}
// SubIdentifierType Setter
// 店铺所在平台
func (r *AlibabaSeakingFeedbackAPIRequest) SetSubIdentifierType(_subIdentifierType string) error {
    r._subIdentifierType = _subIdentifierType
    r.Set("sub_identifier_type", _subIdentifierType)
    return nil
}

// SubIdentifierType Getter
func (r AlibabaSeakingFeedbackAPIRequest) GetSubIdentifierType() string {
    return r._subIdentifierType
}
// Identifier Setter
// erp名称
func (r *AlibabaSeakingFeedbackAPIRequest) SetIdentifier(_identifier string) error {
    r._identifier = _identifier
    r.Set("identifier", _identifier)
    return nil
}

// Identifier Getter
func (r AlibabaSeakingFeedbackAPIRequest) GetIdentifier() string {
    return r._identifier
}
// IdentifierType Setter
// erp用户id
func (r *AlibabaSeakingFeedbackAPIRequest) SetIdentifierType(_identifierType string) error {
    r._identifierType = _identifierType
    r.Set("identifier_type", _identifierType)
    return nil
}

// IdentifierType Getter
func (r AlibabaSeakingFeedbackAPIRequest) GetIdentifierType() string {
    return r._identifierType
}
