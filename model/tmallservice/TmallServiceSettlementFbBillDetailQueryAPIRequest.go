package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicesettlementfbbilldetailqueryAPIRequest 服务商工单结算对账查询-流水查询 API请求
// tmall.service.settlement.fb.bill.detail.query
//
// 服务商工单结算对账查询-流水查询，用于查询服务工单费用流水，含服务费、退款、分成、提现等。
type TmallservicesettlementfbbilldetailqueryAPIRequest struct {
	model.Params
	// 提现开始时间（如果设置了开始时间，需要同时设置截止时间）
	_payTimeStart string
	// 提现截止时间
	_payTimeEnd string
	// 流水截止时间
	_billTimeEnd string
	// 流水开始时间（如果设置了开始时间，需要同时设置截止时间）
	_billTimeStart string
	// 分页大小
	_pageSize int64
	// 分页页码
	_pageNum int64
}

// NewTmallservicesettlementfbbilldetailqueryRequest 初始化TmallservicesettlementfbbilldetailqueryAPIRequest对象
func NewTmallservicesettlementfbbilldetailqueryRequest() *TmallservicesettlementfbbilldetailqueryAPIRequest {
	return &TmallservicesettlementfbbilldetailqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicesettlementfbbilldetailqueryAPIRequest) GetApiMethodName() string {
	return "tmall.service.settlement.fb.bill.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicesettlementfbbilldetailqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicesettlementfbbilldetailqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayTimeStart is PayTimeStart Setter
// 提现开始时间（如果设置了开始时间，需要同时设置截止时间）
func (r *TmallservicesettlementfbbilldetailqueryAPIRequest) SetPayTimeStart(_payTimeStart string) error {
	r._payTimeStart = _payTimeStart
	r.Set("pay_time_start", _payTimeStart)
	return nil
}

// GetPayTimeStart PayTimeStart Getter
func (r TmallservicesettlementfbbilldetailqueryAPIRequest) GetPayTimeStart() string {
	return r._payTimeStart
}

// SetPayTimeEnd is PayTimeEnd Setter
// 提现截止时间
func (r *TmallservicesettlementfbbilldetailqueryAPIRequest) SetPayTimeEnd(_payTimeEnd string) error {
	r._payTimeEnd = _payTimeEnd
	r.Set("pay_time_end", _payTimeEnd)
	return nil
}

// GetPayTimeEnd PayTimeEnd Getter
func (r TmallservicesettlementfbbilldetailqueryAPIRequest) GetPayTimeEnd() string {
	return r._payTimeEnd
}

// SetBillTimeEnd is BillTimeEnd Setter
// 流水截止时间
func (r *TmallservicesettlementfbbilldetailqueryAPIRequest) SetBillTimeEnd(_billTimeEnd string) error {
	r._billTimeEnd = _billTimeEnd
	r.Set("bill_time_end", _billTimeEnd)
	return nil
}

// GetBillTimeEnd BillTimeEnd Getter
func (r TmallservicesettlementfbbilldetailqueryAPIRequest) GetBillTimeEnd() string {
	return r._billTimeEnd
}

// SetBillTimeStart is BillTimeStart Setter
// 流水开始时间（如果设置了开始时间，需要同时设置截止时间）
func (r *TmallservicesettlementfbbilldetailqueryAPIRequest) SetBillTimeStart(_billTimeStart string) error {
	r._billTimeStart = _billTimeStart
	r.Set("bill_time_start", _billTimeStart)
	return nil
}

// GetBillTimeStart BillTimeStart Getter
func (r TmallservicesettlementfbbilldetailqueryAPIRequest) GetBillTimeStart() string {
	return r._billTimeStart
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TmallservicesettlementfbbilldetailqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallservicesettlementfbbilldetailqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 分页页码
func (r *TmallservicesettlementfbbilldetailqueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TmallservicesettlementfbbilldetailqueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}
