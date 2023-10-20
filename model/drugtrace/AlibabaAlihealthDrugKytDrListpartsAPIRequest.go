package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytdrlistpartsAPIRequest 多融查询一个企业的往来单位 API请求
// alibaba.alihealth.drug.kyt.dr.listparts
//
// 查询往来单位列表
type AlibabaalihealthdrugkytdrlistpartsAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 企业名称
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

// NewAlibabaalihealthdrugkytdrlistpartsRequest 初始化AlibabaalihealthdrugkytdrlistpartsAPIRequest对象
func NewAlibabaalihealthdrugkytdrlistpartsRequest() *AlibabaalihealthdrugkytdrlistpartsAPIRequest {
	return &AlibabaalihealthdrugkytdrlistpartsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytdrlistpartsAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.listparts"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytdrlistpartsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytdrlistpartsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthdrugkytdrlistpartsAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytdrlistpartsAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaalihealthdrugkytdrlistpartsAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaalihealthdrugkytdrlistpartsAPIRequest) GetEntName() string {
	return r._entName
}

// SetRefPartnerId is RefPartnerId Setter
// 企业自定义编号
func (r *AlibabaalihealthdrugkytdrlistpartsAPIRequest) SetRefPartnerId(_refPartnerId string) error {
	r._refPartnerId = _refPartnerId
	r.Set("ref_partner_id", _refPartnerId)
	return nil
}

// GetRefPartnerId RefPartnerId Getter
func (r AlibabaalihealthdrugkytdrlistpartsAPIRequest) GetRefPartnerId() string {
	return r._refPartnerId
}

// SetBeginDate is BeginDate Setter
// 开始时间
func (r *AlibabaalihealthdrugkytdrlistpartsAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaalihealthdrugkytdrlistpartsAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *AlibabaalihealthdrugkytdrlistpartsAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaalihealthdrugkytdrlistpartsAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaalihealthdrugkytdrlistpartsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaalihealthdrugkytdrlistpartsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *AlibabaalihealthdrugkytdrlistpartsAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaalihealthdrugkytdrlistpartsAPIRequest) GetPage() int64 {
	return r._page
}
