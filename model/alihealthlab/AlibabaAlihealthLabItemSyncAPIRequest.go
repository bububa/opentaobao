package alihealthlab

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthLabItemSyncAPIRequest 阿里健康检验检测商品发布 API请求
// alibaba.alihealth.lab.item.sync
//
// iSV发布检验检测商品基本信息给健康，内部关联一个淘宝商品或SKU
type AlibabaAlihealthLabItemSyncAPIRequest struct {
	model.Params
	// 检验检测项目isv侧编码
	_isvItemCode string
	// 检验检测项目名称
	_name string
	// 采购价，单位分
	_costPrice int64
	// 性别限制 MALE  FEMALE
	_genderRestriction string
	// 婚否状态限制 MARRIED UNMARRIED
	_maritalStatusRestriction string
	// 额外的属性
	_extraAttributes string
	// 报告产出预计需要的时长，单位毫秒
	_reportGenerationTimeDuration int64
	// 打包的子项目
	_subItems []LabSubItemBrief
	// EFFECTIVE 项目有效， INVALID 项目无效
	_isvItemStatus string
	// 项目关联的门店
	_relatedIsvStoreCodes []string
	// 项目介绍
	_intro string
}

// NewAlibabaAlihealthLabItemSyncRequest 初始化AlibabaAlihealthLabItemSyncAPIRequest对象
func NewAlibabaAlihealthLabItemSyncRequest() *AlibabaAlihealthLabItemSyncAPIRequest {
	return &AlibabaAlihealthLabItemSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthLabItemSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.lab.item.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthLabItemSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is IsvItemCode Setter
// 检验检测项目isv侧编码
func (r *AlibabaAlihealthLabItemSyncAPIRequest) SetIsvItemCode(_isvItemCode string) error {
	r._isvItemCode = _isvItemCode
	r.Set("isv_item_code", _isvItemCode)
	return nil
}

// Get IsvItemCode Getter
func (r AlibabaAlihealthLabItemSyncAPIRequest) GetIsvItemCode() string {
	return r._isvItemCode
}

// Set is Name Setter
// 检验检测项目名称
func (r *AlibabaAlihealthLabItemSyncAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r AlibabaAlihealthLabItemSyncAPIRequest) GetName() string {
	return r._name
}

// Set is CostPrice Setter
// 采购价，单位分
func (r *AlibabaAlihealthLabItemSyncAPIRequest) SetCostPrice(_costPrice int64) error {
	r._costPrice = _costPrice
	r.Set("cost_price", _costPrice)
	return nil
}

// Get CostPrice Getter
func (r AlibabaAlihealthLabItemSyncAPIRequest) GetCostPrice() int64 {
	return r._costPrice
}

// Set is GenderRestriction Setter
// 性别限制 MALE  FEMALE
func (r *AlibabaAlihealthLabItemSyncAPIRequest) SetGenderRestriction(_genderRestriction string) error {
	r._genderRestriction = _genderRestriction
	r.Set("gender_restriction", _genderRestriction)
	return nil
}

// Get GenderRestriction Getter
func (r AlibabaAlihealthLabItemSyncAPIRequest) GetGenderRestriction() string {
	return r._genderRestriction
}

// Set is MaritalStatusRestriction Setter
// 婚否状态限制 MARRIED UNMARRIED
func (r *AlibabaAlihealthLabItemSyncAPIRequest) SetMaritalStatusRestriction(_maritalStatusRestriction string) error {
	r._maritalStatusRestriction = _maritalStatusRestriction
	r.Set("marital_status_restriction", _maritalStatusRestriction)
	return nil
}

// Get MaritalStatusRestriction Getter
func (r AlibabaAlihealthLabItemSyncAPIRequest) GetMaritalStatusRestriction() string {
	return r._maritalStatusRestriction
}

// Set is ExtraAttributes Setter
// 额外的属性
func (r *AlibabaAlihealthLabItemSyncAPIRequest) SetExtraAttributes(_extraAttributes string) error {
	r._extraAttributes = _extraAttributes
	r.Set("extra_attributes", _extraAttributes)
	return nil
}

// Get ExtraAttributes Getter
func (r AlibabaAlihealthLabItemSyncAPIRequest) GetExtraAttributes() string {
	return r._extraAttributes
}

// Set is ReportGenerationTimeDuration Setter
// 报告产出预计需要的时长，单位毫秒
func (r *AlibabaAlihealthLabItemSyncAPIRequest) SetReportGenerationTimeDuration(_reportGenerationTimeDuration int64) error {
	r._reportGenerationTimeDuration = _reportGenerationTimeDuration
	r.Set("report_generation_time_duration", _reportGenerationTimeDuration)
	return nil
}

// Get ReportGenerationTimeDuration Getter
func (r AlibabaAlihealthLabItemSyncAPIRequest) GetReportGenerationTimeDuration() int64 {
	return r._reportGenerationTimeDuration
}

// Set is SubItems Setter
// 打包的子项目
func (r *AlibabaAlihealthLabItemSyncAPIRequest) SetSubItems(_subItems []LabSubItemBrief) error {
	r._subItems = _subItems
	r.Set("sub_items", _subItems)
	return nil
}

// Get SubItems Getter
func (r AlibabaAlihealthLabItemSyncAPIRequest) GetSubItems() []LabSubItemBrief {
	return r._subItems
}

// Set is IsvItemStatus Setter
// EFFECTIVE 项目有效， INVALID 项目无效
func (r *AlibabaAlihealthLabItemSyncAPIRequest) SetIsvItemStatus(_isvItemStatus string) error {
	r._isvItemStatus = _isvItemStatus
	r.Set("isv_item_status", _isvItemStatus)
	return nil
}

// Get IsvItemStatus Getter
func (r AlibabaAlihealthLabItemSyncAPIRequest) GetIsvItemStatus() string {
	return r._isvItemStatus
}

// Set is RelatedIsvStoreCodes Setter
// 项目关联的门店
func (r *AlibabaAlihealthLabItemSyncAPIRequest) SetRelatedIsvStoreCodes(_relatedIsvStoreCodes []string) error {
	r._relatedIsvStoreCodes = _relatedIsvStoreCodes
	r.Set("related_isv_store_codes", _relatedIsvStoreCodes)
	return nil
}

// Get RelatedIsvStoreCodes Getter
func (r AlibabaAlihealthLabItemSyncAPIRequest) GetRelatedIsvStoreCodes() []string {
	return r._relatedIsvStoreCodes
}

// Set is Intro Setter
// 项目介绍
func (r *AlibabaAlihealthLabItemSyncAPIRequest) SetIntro(_intro string) error {
	r._intro = _intro
	r.Set("intro", _intro)
	return nil
}

// Get Intro Getter
func (r AlibabaAlihealthLabItemSyncAPIRequest) GetIntro() string {
	return r._intro
}
