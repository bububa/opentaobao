package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytStorebilllistAPIRequest 零售端平台单据查询 API请求
// alibaba.alihealth.drug.kyt.storebilllist
//
// 零售端平台单据查询
type AlibabaAlihealthDrugKytStorebilllistAPIRequest struct {
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

// NewAlibabaAlihealthDrugKytStorebilllistRequest 初始化AlibabaAlihealthDrugKytStorebilllistAPIRequest对象
func NewAlibabaAlihealthDrugKytStorebilllistRequest() *AlibabaAlihealthDrugKytStorebilllistAPIRequest {
	return &AlibabaAlihealthDrugKytStorebilllistAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytStorebilllistAPIRequest) Reset() {
	r._refEntId = ""
	r._startDate = ""
	r._endDate = ""
	r._billStatus = ""
	r._page = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.storebilllist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytStorebilllistAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetStartDate is StartDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugKytStorebilllistAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugKytStorebilllistAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetBillStatus is BillStatus Setter
// 单据状态(A全部 1上传成功 3处理成功 4处理失败 )
func (r *AlibabaAlihealthDrugKytStorebilllistAPIRequest) SetBillStatus(_billStatus string) error {
	r._billStatus = _billStatus
	r.Set("bill_status", _billStatus)
	return nil
}

// GetBillStatus BillStatus Getter
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetBillStatus() string {
	return r._billStatus
}

// SetPage is Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytStorebilllistAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 页数
func (r *AlibabaAlihealthDrugKytStorebilllistAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaAlihealthDrugKytStorebilllistAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytStorebilllistRequest()
	},
}

// GetAlibabaAlihealthDrugKytStorebilllistRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytStorebilllistAPIRequest
func GetAlibabaAlihealthDrugKytStorebilllistAPIRequest() *AlibabaAlihealthDrugKytStorebilllistAPIRequest {
	return poolAlibabaAlihealthDrugKytStorebilllistAPIRequest.Get().(*AlibabaAlihealthDrugKytStorebilllistAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytStorebilllistAPIRequest 将 AlibabaAlihealthDrugKytStorebilllistAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytStorebilllistAPIRequest(v *AlibabaAlihealthDrugKytStorebilllistAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytStorebilllistAPIRequest.Put(v)
}
