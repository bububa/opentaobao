package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytStorebilllistAPIRequest
零售端平台单据查询 API请求
alibaba.alihealth.drug.kyt.storebilllist

零售端平台单据查询 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.storebilllist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytStorebilllistAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is StartDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugKytStorebilllistAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// Get StartDate Getter
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetStartDate() string {
	return r._startDate
}

// Set is EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugKytStorebilllistAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is BillStatus Setter
// 单据状态(A全部 1上传成功 3处理成功 4处理失败 )
func (r *AlibabaAlihealthDrugKytStorebilllistAPIRequest) SetBillStatus(_billStatus string) error {
	r._billStatus = _billStatus
	r.Set("bill_status", _billStatus)
	return nil
}

// Get BillStatus Getter
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetBillStatus() string {
	return r._billStatus
}

// Set is Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytStorebilllistAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// Get Page Getter
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetPage() int64 {
	return r._page
}

// Set is PageSize Setter
// 页数
func (r *AlibabaAlihealthDrugKytStorebilllistAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaAlihealthDrugKytStorebilllistAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
