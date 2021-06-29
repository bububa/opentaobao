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
    _identifier   string
    // 扩展信息
    _extra   *Extra
    // 语种
    _language   string
    // 商品属性
    _attributes   string
    // 调用来源(erp名称)
    _identifierType   string
    // 标题
    _title   string
    // 商品所在平台（ae/icbu）
    _platform   string
    // 类目id,没有的时候传-1
    _categoryId   int64
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
func (r *AlibabaSeakingAititlegenerateRequest) SetIdentifier(_identifier string) error {
    r._identifier = _identifier
    r.Set("identifier", _identifier)
    return nil
}

// Identifier Getter
func (r AlibabaSeakingAititlegenerateRequest) GetIdentifier() string {
    return r._identifier
}
// Extra Setter
// 扩展信息
func (r *AlibabaSeakingAititlegenerateRequest) SetExtra(_extra *Extra) error {
    r._extra = _extra
    r.Set("extra", _extra)
    return nil
}

// Extra Getter
func (r AlibabaSeakingAititlegenerateRequest) GetExtra() *Extra {
    return r._extra
}
// Language Setter
// 语种
func (r *AlibabaSeakingAititlegenerateRequest) SetLanguage(_language string) error {
    r._language = _language
    r.Set("language", _language)
    return nil
}

// Language Getter
func (r AlibabaSeakingAititlegenerateRequest) GetLanguage() string {
    return r._language
}
// Attributes Setter
// 商品属性
func (r *AlibabaSeakingAititlegenerateRequest) SetAttributes(_attributes string) error {
    r._attributes = _attributes
    r.Set("attributes", _attributes)
    return nil
}

// Attributes Getter
func (r AlibabaSeakingAititlegenerateRequest) GetAttributes() string {
    return r._attributes
}
// IdentifierType Setter
// 调用来源(erp名称)
func (r *AlibabaSeakingAititlegenerateRequest) SetIdentifierType(_identifierType string) error {
    r._identifierType = _identifierType
    r.Set("identifier_type", _identifierType)
    return nil
}

// IdentifierType Getter
func (r AlibabaSeakingAititlegenerateRequest) GetIdentifierType() string {
    return r._identifierType
}
// Title Setter
// 标题
func (r *AlibabaSeakingAititlegenerateRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r AlibabaSeakingAititlegenerateRequest) GetTitle() string {
    return r._title
}
// Platform Setter
// 商品所在平台（ae/icbu）
func (r *AlibabaSeakingAititlegenerateRequest) SetPlatform(_platform string) error {
    r._platform = _platform
    r.Set("platform", _platform)
    return nil
}

// Platform Getter
func (r AlibabaSeakingAititlegenerateRequest) GetPlatform() string {
    return r._platform
}
// CategoryId Setter
// 类目id,没有的时候传-1
func (r *AlibabaSeakingAititlegenerateRequest) SetCategoryId(_categoryId int64) error {
    r._categoryId = _categoryId
    r.Set("category_id", _categoryId)
    return nil
}

// CategoryId Getter
func (r AlibabaSeakingAititlegenerateRequest) GetCategoryId() int64 {
    return r._categoryId
}
