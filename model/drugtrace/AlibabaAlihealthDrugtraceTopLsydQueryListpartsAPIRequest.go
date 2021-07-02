package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest 往来单位查询 API请求
// alibaba.alihealth.drugtrace.top.lsyd.query.listparts
//
// 查询往来单位列表
type AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest struct {
	model.Params
	// 企业唯一标识
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

// NewAlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest 初始化AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydQueryListpartsRequest() *AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest {
	return &AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.query.listparts"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// Get EntName Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) GetEntName() string {
	return r._entName
}

// Set is RefPartnerId Setter
// 企业自定义编号
func (r *AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) SetRefPartnerId(_refPartnerId string) error {
	r._refPartnerId = _refPartnerId
	r.Set("ref_partner_id", _refPartnerId)
	return nil
}

// Get RefPartnerId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) GetRefPartnerId() string {
	return r._refPartnerId
}

// Set is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is Page Setter
// 页码
func (r *AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// Get Page Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) GetPage() int64 {
	return r._page
}

// Set is BeginDate Setter
// 开始时间
func (r *AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// Get BeginDate Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// Set is EndDate Setter
// 结束时间
func (r *AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest) GetEndDate() string {
	return r._endDate
}
