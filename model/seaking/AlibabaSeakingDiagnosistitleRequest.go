package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标题诊断 API请求
alibaba.seaking.diagnosistitle

标题诊断
*/
type AlibabaSeakingDiagnosistitleRequest struct {
    model.Params
    // 类目id,没有的时候传-1
    categoryId   int64
    // 扩展信息
    extra   *Extra
    // erp用户id
    identifier   string
    // 调用来源(erp名称)
    identifierType   string
    // 语种
    language   string
    // 商品所在平台（ae/icbu）
    platform   string
    // 标题
    title   string
}

// 初始化AlibabaSeakingDiagnosistitleRequest对象
func NewAlibabaSeakingDiagnosistitleRequest() *AlibabaSeakingDiagnosistitleRequest{
    return &AlibabaSeakingDiagnosistitleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingDiagnosistitleRequest) GetApiMethodName() string {
    return "alibaba.seaking.diagnosistitle"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingDiagnosistitleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CategoryId Setter
// 类目id,没有的时候传-1
func (r *AlibabaSeakingDiagnosistitleRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

// CategoryId Getter
func (r AlibabaSeakingDiagnosistitleRequest) GetCategoryId() int64 {
    return r.categoryId
}
// Extra Setter
// 扩展信息
func (r *AlibabaSeakingDiagnosistitleRequest) SetExtra(extra *Extra) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

// Extra Getter
func (r AlibabaSeakingDiagnosistitleRequest) GetExtra() *Extra {
    return r.extra
}
// Identifier Setter
// erp用户id
func (r *AlibabaSeakingDiagnosistitleRequest) SetIdentifier(identifier string) error {
    r.identifier = identifier
    r.Set("identifier", identifier)
    return nil
}

// Identifier Getter
func (r AlibabaSeakingDiagnosistitleRequest) GetIdentifier() string {
    return r.identifier
}
// IdentifierType Setter
// 调用来源(erp名称)
func (r *AlibabaSeakingDiagnosistitleRequest) SetIdentifierType(identifierType string) error {
    r.identifierType = identifierType
    r.Set("identifier_type", identifierType)
    return nil
}

// IdentifierType Getter
func (r AlibabaSeakingDiagnosistitleRequest) GetIdentifierType() string {
    return r.identifierType
}
// Language Setter
// 语种
func (r *AlibabaSeakingDiagnosistitleRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

// Language Getter
func (r AlibabaSeakingDiagnosistitleRequest) GetLanguage() string {
    return r.language
}
// Platform Setter
// 商品所在平台（ae/icbu）
func (r *AlibabaSeakingDiagnosistitleRequest) SetPlatform(platform string) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

// Platform Getter
func (r AlibabaSeakingDiagnosistitleRequest) GetPlatform() string {
    return r.platform
}
// Title Setter
// 标题
func (r *AlibabaSeakingDiagnosistitleRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r AlibabaSeakingDiagnosistitleRequest) GetTitle() string {
    return r.title
}
