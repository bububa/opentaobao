package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标题诊断 APIRequest
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

func NewAlibabaSeakingDiagnosistitleRequest() *AlibabaSeakingDiagnosistitleRequest{
    return &AlibabaSeakingDiagnosistitleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSeakingDiagnosistitleRequest) GetApiMethodName() string {
    return "alibaba.seaking.diagnosistitle"
}

func (r AlibabaSeakingDiagnosistitleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSeakingDiagnosistitleRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r AlibabaSeakingDiagnosistitleRequest) GetCategoryId() int64 {
    return r.categoryId
}

func (r *AlibabaSeakingDiagnosistitleRequest) SetExtra(extra *Extra) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

func (r AlibabaSeakingDiagnosistitleRequest) GetExtra() *Extra {
    return r.extra
}

func (r *AlibabaSeakingDiagnosistitleRequest) SetIdentifier(identifier string) error {
    r.identifier = identifier
    r.Set("identifier", identifier)
    return nil
}

func (r AlibabaSeakingDiagnosistitleRequest) GetIdentifier() string {
    return r.identifier
}

func (r *AlibabaSeakingDiagnosistitleRequest) SetIdentifierType(identifierType string) error {
    r.identifierType = identifierType
    r.Set("identifier_type", identifierType)
    return nil
}

func (r AlibabaSeakingDiagnosistitleRequest) GetIdentifierType() string {
    return r.identifierType
}

func (r *AlibabaSeakingDiagnosistitleRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

func (r AlibabaSeakingDiagnosistitleRequest) GetLanguage() string {
    return r.language
}

func (r *AlibabaSeakingDiagnosistitleRequest) SetPlatform(platform string) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

func (r AlibabaSeakingDiagnosistitleRequest) GetPlatform() string {
    return r.platform
}

func (r *AlibabaSeakingDiagnosistitleRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r AlibabaSeakingDiagnosistitleRequest) GetTitle() string {
    return r.title
}

