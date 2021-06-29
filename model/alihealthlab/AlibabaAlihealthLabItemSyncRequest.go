package alihealthlab

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康检验检测商品发布 API请求
alibaba.alihealth.lab.item.sync

iSV发布检验检测商品基本信息给健康，内部关联一个淘宝商品或SKU
*/
type AlibabaAlihealthLabItemSyncRequest struct {
    model.Params
    // 检验检测项目isv侧编码
    isvItemCode   string
    // 检验检测项目名称
    name   string
    // 采购价，单位分
    costPrice   int64
    // 性别限制 MALE  FEMALE
    genderRestriction   string
    // 婚否状态限制 MARRIED UNMARRIED
    maritalStatusRestriction   string
    // 额外的属性
    extraAttributes   string
    // 报告产出预计需要的时长，单位毫秒
    reportGenerationTimeDuration   int64
    // 打包的子项目
    subItems   []LabSubItemBrief
    // EFFECTIVE 项目有效， INVALID 项目无效
    isvItemStatus   string
    // 项目关联的门店
    relatedIsvStoreCodes   []string
    // 项目介绍
    intro   string
}

// 初始化AlibabaAlihealthLabItemSyncRequest对象
func NewAlibabaAlihealthLabItemSyncRequest() *AlibabaAlihealthLabItemSyncRequest{
    return &AlibabaAlihealthLabItemSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthLabItemSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.lab.item.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthLabItemSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsvItemCode Setter
// 检验检测项目isv侧编码
func (r *AlibabaAlihealthLabItemSyncRequest) SetIsvItemCode(isvItemCode string) error {
    r.isvItemCode = isvItemCode
    r.Set("isv_item_code", isvItemCode)
    return nil
}

// IsvItemCode Getter
func (r AlibabaAlihealthLabItemSyncRequest) GetIsvItemCode() string {
    return r.isvItemCode
}
// Name Setter
// 检验检测项目名称
func (r *AlibabaAlihealthLabItemSyncRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlibabaAlihealthLabItemSyncRequest) GetName() string {
    return r.name
}
// CostPrice Setter
// 采购价，单位分
func (r *AlibabaAlihealthLabItemSyncRequest) SetCostPrice(costPrice int64) error {
    r.costPrice = costPrice
    r.Set("cost_price", costPrice)
    return nil
}

// CostPrice Getter
func (r AlibabaAlihealthLabItemSyncRequest) GetCostPrice() int64 {
    return r.costPrice
}
// GenderRestriction Setter
// 性别限制 MALE  FEMALE
func (r *AlibabaAlihealthLabItemSyncRequest) SetGenderRestriction(genderRestriction string) error {
    r.genderRestriction = genderRestriction
    r.Set("gender_restriction", genderRestriction)
    return nil
}

// GenderRestriction Getter
func (r AlibabaAlihealthLabItemSyncRequest) GetGenderRestriction() string {
    return r.genderRestriction
}
// MaritalStatusRestriction Setter
// 婚否状态限制 MARRIED UNMARRIED
func (r *AlibabaAlihealthLabItemSyncRequest) SetMaritalStatusRestriction(maritalStatusRestriction string) error {
    r.maritalStatusRestriction = maritalStatusRestriction
    r.Set("marital_status_restriction", maritalStatusRestriction)
    return nil
}

// MaritalStatusRestriction Getter
func (r AlibabaAlihealthLabItemSyncRequest) GetMaritalStatusRestriction() string {
    return r.maritalStatusRestriction
}
// ExtraAttributes Setter
// 额外的属性
func (r *AlibabaAlihealthLabItemSyncRequest) SetExtraAttributes(extraAttributes string) error {
    r.extraAttributes = extraAttributes
    r.Set("extra_attributes", extraAttributes)
    return nil
}

// ExtraAttributes Getter
func (r AlibabaAlihealthLabItemSyncRequest) GetExtraAttributes() string {
    return r.extraAttributes
}
// ReportGenerationTimeDuration Setter
// 报告产出预计需要的时长，单位毫秒
func (r *AlibabaAlihealthLabItemSyncRequest) SetReportGenerationTimeDuration(reportGenerationTimeDuration int64) error {
    r.reportGenerationTimeDuration = reportGenerationTimeDuration
    r.Set("report_generation_time_duration", reportGenerationTimeDuration)
    return nil
}

// ReportGenerationTimeDuration Getter
func (r AlibabaAlihealthLabItemSyncRequest) GetReportGenerationTimeDuration() int64 {
    return r.reportGenerationTimeDuration
}
// SubItems Setter
// 打包的子项目
func (r *AlibabaAlihealthLabItemSyncRequest) SetSubItems(subItems []LabSubItemBrief) error {
    r.subItems = subItems
    r.Set("sub_items", subItems)
    return nil
}

// SubItems Getter
func (r AlibabaAlihealthLabItemSyncRequest) GetSubItems() []LabSubItemBrief {
    return r.subItems
}
// IsvItemStatus Setter
// EFFECTIVE 项目有效， INVALID 项目无效
func (r *AlibabaAlihealthLabItemSyncRequest) SetIsvItemStatus(isvItemStatus string) error {
    r.isvItemStatus = isvItemStatus
    r.Set("isv_item_status", isvItemStatus)
    return nil
}

// IsvItemStatus Getter
func (r AlibabaAlihealthLabItemSyncRequest) GetIsvItemStatus() string {
    return r.isvItemStatus
}
// RelatedIsvStoreCodes Setter
// 项目关联的门店
func (r *AlibabaAlihealthLabItemSyncRequest) SetRelatedIsvStoreCodes(relatedIsvStoreCodes []string) error {
    r.relatedIsvStoreCodes = relatedIsvStoreCodes
    r.Set("related_isv_store_codes", relatedIsvStoreCodes)
    return nil
}

// RelatedIsvStoreCodes Getter
func (r AlibabaAlihealthLabItemSyncRequest) GetRelatedIsvStoreCodes() []string {
    return r.relatedIsvStoreCodes
}
// Intro Setter
// 项目介绍
func (r *AlibabaAlihealthLabItemSyncRequest) SetIntro(intro string) error {
    r.intro = intro
    r.Set("intro", intro)
    return nil
}

// Intro Getter
func (r AlibabaAlihealthLabItemSyncRequest) GetIntro() string {
    return r.intro
}
