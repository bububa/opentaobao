package alihealthlab

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthlabitemsyncAPIRequest 阿里健康检验检测商品发布 API请求
// alibaba.alihealth.lab.item.sync
//
// iSV发布检验检测商品基本信息给健康，内部关联一个淘宝商品或SKU
type AlibabaalihealthlabitemsyncAPIRequest struct {
	model.Params
	// 打包的子项目
	_subItems []LabSubItemBrief
	// 项目关联的门店
	_relatedIsvStoreCodes []string
	// 检验检测项目isv侧编码
	_isvItemCode string
	// 检验检测项目名称
	_name string
	// 性别限制 MALE  FEMALE
	_genderRestriction string
	// 婚否状态限制 MARRIED UNMARRIED
	_maritalStatusRestriction string
	// 额外的属性
	_extraAttributes string
	// EFFECTIVE 项目有效， INVALID 项目无效
	_isvItemStatus string
	// 项目介绍
	_intro string
	// 采购价，单位分
	_costPrice int64
	// 报告产出预计需要的时长，单位毫秒
	_reportGenerationTimeDuration int64
}

// NewAlibabaalihealthlabitemsyncRequest 初始化AlibabaalihealthlabitemsyncAPIRequest对象
func NewAlibabaalihealthlabitemsyncRequest() *AlibabaalihealthlabitemsyncAPIRequest {
	return &AlibabaalihealthlabitemsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthlabitemsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.lab.item.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthlabitemsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthlabitemsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubItems is SubItems Setter
// 打包的子项目
func (r *AlibabaalihealthlabitemsyncAPIRequest) SetSubItems(_subItems []LabSubItemBrief) error {
	r._subItems = _subItems
	r.Set("sub_items", _subItems)
	return nil
}

// GetSubItems SubItems Getter
func (r AlibabaalihealthlabitemsyncAPIRequest) GetSubItems() []LabSubItemBrief {
	return r._subItems
}

// SetRelatedIsvStoreCodes is RelatedIsvStoreCodes Setter
// 项目关联的门店
func (r *AlibabaalihealthlabitemsyncAPIRequest) SetRelatedIsvStoreCodes(_relatedIsvStoreCodes []string) error {
	r._relatedIsvStoreCodes = _relatedIsvStoreCodes
	r.Set("related_isv_store_codes", _relatedIsvStoreCodes)
	return nil
}

// GetRelatedIsvStoreCodes RelatedIsvStoreCodes Getter
func (r AlibabaalihealthlabitemsyncAPIRequest) GetRelatedIsvStoreCodes() []string {
	return r._relatedIsvStoreCodes
}

// SetIsvItemCode is IsvItemCode Setter
// 检验检测项目isv侧编码
func (r *AlibabaalihealthlabitemsyncAPIRequest) SetIsvItemCode(_isvItemCode string) error {
	r._isvItemCode = _isvItemCode
	r.Set("isv_item_code", _isvItemCode)
	return nil
}

// GetIsvItemCode IsvItemCode Getter
func (r AlibabaalihealthlabitemsyncAPIRequest) GetIsvItemCode() string {
	return r._isvItemCode
}

// SetName is Name Setter
// 检验检测项目名称
func (r *AlibabaalihealthlabitemsyncAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabaalihealthlabitemsyncAPIRequest) GetName() string {
	return r._name
}

// SetGenderRestriction is GenderRestriction Setter
// 性别限制 MALE  FEMALE
func (r *AlibabaalihealthlabitemsyncAPIRequest) SetGenderRestriction(_genderRestriction string) error {
	r._genderRestriction = _genderRestriction
	r.Set("gender_restriction", _genderRestriction)
	return nil
}

// GetGenderRestriction GenderRestriction Getter
func (r AlibabaalihealthlabitemsyncAPIRequest) GetGenderRestriction() string {
	return r._genderRestriction
}

// SetMaritalStatusRestriction is MaritalStatusRestriction Setter
// 婚否状态限制 MARRIED UNMARRIED
func (r *AlibabaalihealthlabitemsyncAPIRequest) SetMaritalStatusRestriction(_maritalStatusRestriction string) error {
	r._maritalStatusRestriction = _maritalStatusRestriction
	r.Set("marital_status_restriction", _maritalStatusRestriction)
	return nil
}

// GetMaritalStatusRestriction MaritalStatusRestriction Getter
func (r AlibabaalihealthlabitemsyncAPIRequest) GetMaritalStatusRestriction() string {
	return r._maritalStatusRestriction
}

// SetExtraAttributes is ExtraAttributes Setter
// 额外的属性
func (r *AlibabaalihealthlabitemsyncAPIRequest) SetExtraAttributes(_extraAttributes string) error {
	r._extraAttributes = _extraAttributes
	r.Set("extra_attributes", _extraAttributes)
	return nil
}

// GetExtraAttributes ExtraAttributes Getter
func (r AlibabaalihealthlabitemsyncAPIRequest) GetExtraAttributes() string {
	return r._extraAttributes
}

// SetIsvItemStatus is IsvItemStatus Setter
// EFFECTIVE 项目有效， INVALID 项目无效
func (r *AlibabaalihealthlabitemsyncAPIRequest) SetIsvItemStatus(_isvItemStatus string) error {
	r._isvItemStatus = _isvItemStatus
	r.Set("isv_item_status", _isvItemStatus)
	return nil
}

// GetIsvItemStatus IsvItemStatus Getter
func (r AlibabaalihealthlabitemsyncAPIRequest) GetIsvItemStatus() string {
	return r._isvItemStatus
}

// SetIntro is Intro Setter
// 项目介绍
func (r *AlibabaalihealthlabitemsyncAPIRequest) SetIntro(_intro string) error {
	r._intro = _intro
	r.Set("intro", _intro)
	return nil
}

// GetIntro Intro Getter
func (r AlibabaalihealthlabitemsyncAPIRequest) GetIntro() string {
	return r._intro
}

// SetCostPrice is CostPrice Setter
// 采购价，单位分
func (r *AlibabaalihealthlabitemsyncAPIRequest) SetCostPrice(_costPrice int64) error {
	r._costPrice = _costPrice
	r.Set("cost_price", _costPrice)
	return nil
}

// GetCostPrice CostPrice Getter
func (r AlibabaalihealthlabitemsyncAPIRequest) GetCostPrice() int64 {
	return r._costPrice
}

// SetReportGenerationTimeDuration is ReportGenerationTimeDuration Setter
// 报告产出预计需要的时长，单位毫秒
func (r *AlibabaalihealthlabitemsyncAPIRequest) SetReportGenerationTimeDuration(_reportGenerationTimeDuration int64) error {
	r._reportGenerationTimeDuration = _reportGenerationTimeDuration
	r.Set("report_generation_time_duration", _reportGenerationTimeDuration)
	return nil
}

// GetReportGenerationTimeDuration ReportGenerationTimeDuration Getter
func (r AlibabaalihealthlabitemsyncAPIRequest) GetReportGenerationTimeDuration() int64 {
	return r._reportGenerationTimeDuration
}
