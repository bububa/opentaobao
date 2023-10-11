package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest 物流企业查询货主企业信息 API请求
// alibaba.alihealth.drug.kyt.wes.synonymauths
//
// 物流企业查询货主企业信息
type AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 服务时间
	_licenseToken string
	// 企业名称
	_entName string
	// 货主自定义编号
	_synOwnEntId string
	// 页码
	_pageSize int64
	// 页面大小
	_page int64
}

// NewAlibabaAlihealthDrugKytWesSynonymauthsRequest 初始化AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest对象
func NewAlibabaAlihealthDrugKytWesSynonymauthsRequest() *AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest {
	return &AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.synonymauths"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) GetEntName() string {
	return r._entName
}

// SetSynOwnEntId is SynOwnEntId Setter
// 货主自定义编号
func (r *AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) SetSynOwnEntId(_synOwnEntId string) error {
	r._synOwnEntId = _synOwnEntId
	r.Set("syn_own_ent_id", _synOwnEntId)
	return nil
}

// GetSynOwnEntId SynOwnEntId Getter
func (r AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) GetSynOwnEntId() string {
	return r._synOwnEntId
}

// SetPageSize is PageSize Setter
// 页码
func (r *AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页面大小
func (r *AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytWesSynonymauthsAPIRequest) GetPage() int64 {
	return r._page
}
