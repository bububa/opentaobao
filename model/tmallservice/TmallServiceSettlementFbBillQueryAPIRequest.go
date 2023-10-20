package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicesettlementfbbillqueryAPIRequest 服务商工单结算对账查询 API请求
// tmall.service.settlement.fb.bill.query
//
// 服务商工单结算对账查询，用于查询服务工单对应的结算费用情况，含工单对应的服务费、退款、增加费用、分成费用、提现流水
type TmallservicesettlementfbbillqueryAPIRequest struct {
	model.Params
	// 提现开始时间（与截止时间成对使用）
	_payTimeStart string
	// 提现截止时间
	_payTimeEnd string
	// 流水截止时间
	_billTimeEnd string
	// 流水开始时间（与截止时间配对使用）
	_billTimeStart string
	// 分页大小
	_pageSize int64
	// 分页页码
	_pageNum int64
}

// NewTmallservicesettlementfbbillqueryRequest 初始化TmallservicesettlementfbbillqueryAPIRequest对象
func NewTmallservicesettlementfbbillqueryRequest() *TmallservicesettlementfbbillqueryAPIRequest {
	return &TmallservicesettlementfbbillqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicesettlementfbbillqueryAPIRequest) GetApiMethodName() string {
	return "tmall.service.settlement.fb.bill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicesettlementfbbillqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicesettlementfbbillqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayTimeStart is PayTimeStart Setter
// 提现开始时间（与截止时间成对使用）
func (r *TmallservicesettlementfbbillqueryAPIRequest) SetPayTimeStart(_payTimeStart string) error {
	r._payTimeStart = _payTimeStart
	r.Set("pay_time_start", _payTimeStart)
	return nil
}

// GetPayTimeStart PayTimeStart Getter
func (r TmallservicesettlementfbbillqueryAPIRequest) GetPayTimeStart() string {
	return r._payTimeStart
}

// SetPayTimeEnd is PayTimeEnd Setter
// 提现截止时间
func (r *TmallservicesettlementfbbillqueryAPIRequest) SetPayTimeEnd(_payTimeEnd string) error {
	r._payTimeEnd = _payTimeEnd
	r.Set("pay_time_end", _payTimeEnd)
	return nil
}

// GetPayTimeEnd PayTimeEnd Getter
func (r TmallservicesettlementfbbillqueryAPIRequest) GetPayTimeEnd() string {
	return r._payTimeEnd
}

// SetBillTimeEnd is BillTimeEnd Setter
// 流水截止时间
func (r *TmallservicesettlementfbbillqueryAPIRequest) SetBillTimeEnd(_billTimeEnd string) error {
	r._billTimeEnd = _billTimeEnd
	r.Set("bill_time_end", _billTimeEnd)
	return nil
}

// GetBillTimeEnd BillTimeEnd Getter
func (r TmallservicesettlementfbbillqueryAPIRequest) GetBillTimeEnd() string {
	return r._billTimeEnd
}

// SetBillTimeStart is BillTimeStart Setter
// 流水开始时间（与截止时间配对使用）
func (r *TmallservicesettlementfbbillqueryAPIRequest) SetBillTimeStart(_billTimeStart string) error {
	r._billTimeStart = _billTimeStart
	r.Set("bill_time_start", _billTimeStart)
	return nil
}

// GetBillTimeStart BillTimeStart Getter
func (r TmallservicesettlementfbbillqueryAPIRequest) GetBillTimeStart() string {
	return r._billTimeStart
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TmallservicesettlementfbbillqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallservicesettlementfbbillqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 分页页码
func (r *TmallservicesettlementfbbillqueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TmallservicesettlementfbbillqueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}
