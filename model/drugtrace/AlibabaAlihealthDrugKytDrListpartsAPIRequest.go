package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrListpartsAPIRequest
多融查询一个企业的往来单位 API请求
alibaba.alihealth.drug.kyt.dr.listparts

查询往来单位列表 */
type AlibabaAlihealthDrugKytDrListpartsAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 企业名称
	_entName string
	// 企业自定义编号
	_refPartnerId string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
	// 开始时间
	_beginDate string
	// 结束时间
	_endDate string
}

// NewAlibabaAlihealthDrugKytDrListpartsRequest 初始化AlibabaAlihealthDrugKytDrListpartsAPIRequest对象
func NewAlibabaAlihealthDrugKytDrListpartsRequest() *AlibabaAlihealthDrugKytDrListpartsAPIRequest {
	return &AlibabaAlihealthDrugKytDrListpartsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrListpartsAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.listparts"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrListpartsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytDrListpartsAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugKytDrListpartsAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytDrListpartsAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// Get EntName Getter
func (r AlibabaAlihealthDrugKytDrListpartsAPIRequest) GetEntName() string {
	return r._entName
}

// Set is RefPartnerId Setter
// 企业自定义编号
func (r *AlibabaAlihealthDrugKytDrListpartsAPIRequest) SetRefPartnerId(_refPartnerId string) error {
	r._refPartnerId = _refPartnerId
	r.Set("ref_partner_id", _refPartnerId)
	return nil
}

// Get RefPartnerId Getter
func (r AlibabaAlihealthDrugKytDrListpartsAPIRequest) GetRefPartnerId() string {
	return r._refPartnerId
}

// Set is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytDrListpartsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaAlihealthDrugKytDrListpartsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytDrListpartsAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// Get Page Getter
func (r AlibabaAlihealthDrugKytDrListpartsAPIRequest) GetPage() int64 {
	return r._page
}

// Set is BeginDate Setter
// 开始时间
func (r *AlibabaAlihealthDrugKytDrListpartsAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// Get BeginDate Getter
func (r AlibabaAlihealthDrugKytDrListpartsAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// Set is EndDate Setter
// 结束时间
func (r *AlibabaAlihealthDrugKytDrListpartsAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r AlibabaAlihealthDrugKytDrListpartsAPIRequest) GetEndDate() string {
	return r._endDate
}
