package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标题智能优化 API请求
alibaba.seaking.aititlegenerate

标题智能优化
*/
type AlibabaSeakingAititlegenerateRequest struct {
    model.Params
    // erp用户id
    identifier   string
    // 扩展信息
    extra   *Extra
    // 语种
    language   string
    // 商品属性
    attributes   string
    // 调用来源(erp名称)
    identifierType   string
    // 标题
    title   string
    // 商品所在平台（ae/icbu）
    platform   string
    // 类目id,没有的时候传-1
    categoryId   int64
}

// 初始化AlibabaSeakingAititlegenerateRequest对象
func NewAlibabaSeakingAititlegenerateRequest() *AlibabaSeakingAititlegenerateRequest{
    return &AlibabaSeakingAititlegenerateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingAititlegenerateRequest) GetApiMethodName() string {
    return "alibaba.seaking.aititlegenerate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingAititlegenerateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Identifier Setter
// erp用户id
func (r *AlibabaSeakingAititlegenerateRequest) SetIdentifier(identifier string) error {
    r.identifier = identifier
    r.Set("identifier", identifier)
    return nil
}

// Identifier Getter
func (r AlibabaSeakingAititlegenerateRequest) GetIdentifier() string {
    return r.identifier
}
// Extra Setter
// 扩展信息
func (r *AlibabaSeakingAititlegenerateRequest) SetExtra(extra *Extra) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

// Extra Getter
func (r AlibabaSeakingAititlegenerateRequest) GetExtra() *Extra {
    return r.extra
}
// Language Setter
// 语种
func (r *AlibabaSeakingAititlegenerateRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

// Language Getter
func (r AlibabaSeakingAititlegenerateRequest) GetLanguage() string {
    return r.language
}
// Attributes Setter
// 商品属性
func (r *AlibabaSeakingAititlegenerateRequest) SetAttributes(attributes string) error {
    r.attributes = attributes
    r.Set("attributes", attributes)
    return nil
}

// Attributes Getter
func (r AlibabaSeakingAititlegenerateRequest) GetAttributes() string {
    return r.attributes
}
// IdentifierType Setter
// 调用来源(erp名称)
func (r *AlibabaSeakingAititlegenerateRequest) SetIdentifierType(identifierType string) error {
    r.identifierType = identifierType
    r.Set("identifier_type", identifierType)
    return nil
}

// IdentifierType Getter
func (r AlibabaSeakingAititlegenerateRequest) GetIdentifierType() string {
    return r.identifierType
}
// Title Setter
// 标题
func (r *AlibabaSeakingAititlegenerateRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r AlibabaSeakingAititlegenerateRequest) GetTitle() string {
    return r.title
}
// Platform Setter
// 商品所在平台（ae/icbu）
func (r *AlibabaSeakingAititlegenerateRequest) SetPlatform(platform string) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

// Platform Getter
func (r AlibabaSeakingAititlegenerateRequest) GetPlatform() string {
    return r.platform
}
// CategoryId Setter
// 类目id,没有的时候传-1
func (r *AlibabaSeakingAititlegenerateRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

// CategoryId Getter
func (r AlibabaSeakingAititlegenerateRequest) GetCategoryId() int64 {
    return r.categoryId
}
