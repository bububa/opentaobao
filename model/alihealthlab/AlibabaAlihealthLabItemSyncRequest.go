package alihealthlab

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康检验检测商品发布 APIRequest
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

func NewAlibabaAlihealthLabItemSyncRequest() *AlibabaAlihealthLabItemSyncRequest{
    return &AlibabaAlihealthLabItemSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthLabItemSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.lab.item.sync"
}

func (r AlibabaAlihealthLabItemSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthLabItemSyncRequest) SetIsvItemCode(isvItemCode string) error {
    r.isvItemCode = isvItemCode
    r.Set("isv_item_code", isvItemCode)
    return nil
}

func (r AlibabaAlihealthLabItemSyncRequest) GetIsvItemCode() string {
    return r.isvItemCode
}

func (r *AlibabaAlihealthLabItemSyncRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlibabaAlihealthLabItemSyncRequest) GetName() string {
    return r.name
}

func (r *AlibabaAlihealthLabItemSyncRequest) SetCostPrice(costPrice int64) error {
    r.costPrice = costPrice
    r.Set("cost_price", costPrice)
    return nil
}

func (r AlibabaAlihealthLabItemSyncRequest) GetCostPrice() int64 {
    return r.costPrice
}

func (r *AlibabaAlihealthLabItemSyncRequest) SetGenderRestriction(genderRestriction string) error {
    r.genderRestriction = genderRestriction
    r.Set("gender_restriction", genderRestriction)
    return nil
}

func (r AlibabaAlihealthLabItemSyncRequest) GetGenderRestriction() string {
    return r.genderRestriction
}

func (r *AlibabaAlihealthLabItemSyncRequest) SetMaritalStatusRestriction(maritalStatusRestriction string) error {
    r.maritalStatusRestriction = maritalStatusRestriction
    r.Set("marital_status_restriction", maritalStatusRestriction)
    return nil
}

func (r AlibabaAlihealthLabItemSyncRequest) GetMaritalStatusRestriction() string {
    return r.maritalStatusRestriction
}

func (r *AlibabaAlihealthLabItemSyncRequest) SetExtraAttributes(extraAttributes string) error {
    r.extraAttributes = extraAttributes
    r.Set("extra_attributes", extraAttributes)
    return nil
}

func (r AlibabaAlihealthLabItemSyncRequest) GetExtraAttributes() string {
    return r.extraAttributes
}

func (r *AlibabaAlihealthLabItemSyncRequest) SetReportGenerationTimeDuration(reportGenerationTimeDuration int64) error {
    r.reportGenerationTimeDuration = reportGenerationTimeDuration
    r.Set("report_generation_time_duration", reportGenerationTimeDuration)
    return nil
}

func (r AlibabaAlihealthLabItemSyncRequest) GetReportGenerationTimeDuration() int64 {
    return r.reportGenerationTimeDuration
}

func (r *AlibabaAlihealthLabItemSyncRequest) SetSubItems(subItems []LabSubItemBrief) error {
    r.subItems = subItems
    r.Set("sub_items", subItems)
    return nil
}

func (r AlibabaAlihealthLabItemSyncRequest) GetSubItems() []LabSubItemBrief {
    return r.subItems
}

func (r *AlibabaAlihealthLabItemSyncRequest) SetIsvItemStatus(isvItemStatus string) error {
    r.isvItemStatus = isvItemStatus
    r.Set("isv_item_status", isvItemStatus)
    return nil
}

func (r AlibabaAlihealthLabItemSyncRequest) GetIsvItemStatus() string {
    return r.isvItemStatus
}

func (r *AlibabaAlihealthLabItemSyncRequest) SetRelatedIsvStoreCodes(relatedIsvStoreCodes []string) error {
    r.relatedIsvStoreCodes = relatedIsvStoreCodes
    r.Set("related_isv_store_codes", relatedIsvStoreCodes)
    return nil
}

func (r AlibabaAlihealthLabItemSyncRequest) GetRelatedIsvStoreCodes() []string {
    return r.relatedIsvStoreCodes
}

func (r *AlibabaAlihealthLabItemSyncRequest) SetIntro(intro string) error {
    r.intro = intro
    r.Set("intro", intro)
    return nil
}

func (r AlibabaAlihealthLabItemSyncRequest) GetIntro() string {
    return r.intro
}

