package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesListpartsAPIRequest 查询往来单位列表 API请求
// alibaba.alihealth.drug.kyt.wes.listparts
//
// 查询往来单位列表
type AlibabaAlihealthDrugKytWesListpartsAPIRequest struct {
	model.Params
	// 企业唯一标识
	_refEntId string
	// 服务时间
	_licenseToken string
	// 测试企业
	_entName string
	// 企业自定义编号
	_refPartnerId string
	// 开始时间
	_beginDate string
	// 结束时间
	_endDate string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
}

// NewAlibabaAlihealthDrugKytWesListpartsRequest 初始化AlibabaAlihealthDrugKytWesListpartsAPIRequest对象
func NewAlibabaAlihealthDrugKytWesListpartsRequest() *AlibabaAlihealthDrugKytWesListpartsAPIRequest {
	return &AlibabaAlihealthDrugKytWesListpartsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytWesListpartsAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.listparts"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytWesListpartsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytWesListpartsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugKytWesListpartsAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytWesListpartsAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaAlihealthDrugKytWesListpartsAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugKytWesListpartsAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetEntName is EntName Setter
// 测试企业
func (r *AlibabaAlihealthDrugKytWesListpartsAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytWesListpartsAPIRequest) GetEntName() string {
	return r._entName
}

// SetRefPartnerId is RefPartnerId Setter
// 企业自定义编号
func (r *AlibabaAlihealthDrugKytWesListpartsAPIRequest) SetRefPartnerId(_refPartnerId string) error {
	r._refPartnerId = _refPartnerId
	r.Set("ref_partner_id", _refPartnerId)
	return nil
}

// GetRefPartnerId RefPartnerId Getter
func (r AlibabaAlihealthDrugKytWesListpartsAPIRequest) GetRefPartnerId() string {
	return r._refPartnerId
}

// SetBeginDate is BeginDate Setter
// 开始时间
func (r *AlibabaAlihealthDrugKytWesListpartsAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaAlihealthDrugKytWesListpartsAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *AlibabaAlihealthDrugKytWesListpartsAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaAlihealthDrugKytWesListpartsAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytWesListpartsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytWesListpartsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytWesListpartsAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytWesListpartsAPIRequest) GetPage() int64 {
	return r._page
}
