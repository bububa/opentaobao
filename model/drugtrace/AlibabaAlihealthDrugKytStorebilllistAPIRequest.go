package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytstorebilllistAPIRequest 零售端平台单据查询 API请求
// alibaba.alihealth.drug.kyt.storebilllist
//
// 零售端平台单据查询
type AlibabaalihealthdrugkytstorebilllistAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 开始日期
	_startDate string
	// 结束日期
	_endDate string
	// 单据状态(A全部 1上传成功 3处理成功 4处理失败 )
	_billStatus string
	// 页码
	_page int64
	// 页数
	_pageSize int64
}

// NewAlibabaalihealthdrugkytstorebilllistRequest 初始化AlibabaalihealthdrugkytstorebilllistAPIRequest对象
func NewAlibabaalihealthdrugkytstorebilllistRequest() *AlibabaalihealthdrugkytstorebilllistAPIRequest {
	return &AlibabaalihealthdrugkytstorebilllistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytstorebilllistAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.storebilllist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytstorebilllistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytstorebilllistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthdrugkytstorebilllistAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytstorebilllistAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetStartDate is StartDate Setter
// 开始日期
func (r *AlibabaalihealthdrugkytstorebilllistAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r AlibabaalihealthdrugkytstorebilllistAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束日期
func (r *AlibabaalihealthdrugkytstorebilllistAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaalihealthdrugkytstorebilllistAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetBillStatus is BillStatus Setter
// 单据状态(A全部 1上传成功 3处理成功 4处理失败 )
func (r *AlibabaalihealthdrugkytstorebilllistAPIRequest) SetBillStatus(_billStatus string) error {
	r._billStatus = _billStatus
	r.Set("bill_status", _billStatus)
	return nil
}

// GetBillStatus BillStatus Getter
func (r AlibabaalihealthdrugkytstorebilllistAPIRequest) GetBillStatus() string {
	return r._billStatus
}

// SetPage is Page Setter
// 页码
func (r *AlibabaalihealthdrugkytstorebilllistAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaalihealthdrugkytstorebilllistAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 页数
func (r *AlibabaalihealthdrugkytstorebilllistAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaalihealthdrugkytstorebilllistAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
