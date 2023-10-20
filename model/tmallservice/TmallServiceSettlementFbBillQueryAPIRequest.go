package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettlementFbBillQueryAPIRequest 服务商工单结算对账查询 API请求
// tmall.service.settlement.fb.bill.query
//
// 服务商工单结算对账查询，用于查询服务工单对应的结算费用情况，含工单对应的服务费、退款、增加费用、分成费用、提现流水
type TmallServiceSettlementFbBillQueryAPIRequest struct {
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

// NewTmallServiceSettlementFbBillQueryRequest 初始化TmallServiceSettlementFbBillQueryAPIRequest对象
func NewTmallServiceSettlementFbBillQueryRequest() *TmallServiceSettlementFbBillQueryAPIRequest {
	return &TmallServiceSettlementFbBillQueryAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServiceSettlementFbBillQueryAPIRequest) Reset() {
	r._payTimeStart = ""
	r._payTimeEnd = ""
	r._billTimeEnd = ""
	r._billTimeStart = ""
	r._pageSize = 0
	r._pageNum = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServiceSettlementFbBillQueryAPIRequest) GetApiMethodName() string {
	return "tmall.service.settlement.fb.bill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServiceSettlementFbBillQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServiceSettlementFbBillQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayTimeStart is PayTimeStart Setter
// 提现开始时间（与截止时间成对使用）
func (r *TmallServiceSettlementFbBillQueryAPIRequest) SetPayTimeStart(_payTimeStart string) error {
	r._payTimeStart = _payTimeStart
	r.Set("pay_time_start", _payTimeStart)
	return nil
}

// GetPayTimeStart PayTimeStart Getter
func (r TmallServiceSettlementFbBillQueryAPIRequest) GetPayTimeStart() string {
	return r._payTimeStart
}

// SetPayTimeEnd is PayTimeEnd Setter
// 提现截止时间
func (r *TmallServiceSettlementFbBillQueryAPIRequest) SetPayTimeEnd(_payTimeEnd string) error {
	r._payTimeEnd = _payTimeEnd
	r.Set("pay_time_end", _payTimeEnd)
	return nil
}

// GetPayTimeEnd PayTimeEnd Getter
func (r TmallServiceSettlementFbBillQueryAPIRequest) GetPayTimeEnd() string {
	return r._payTimeEnd
}

// SetBillTimeEnd is BillTimeEnd Setter
// 流水截止时间
func (r *TmallServiceSettlementFbBillQueryAPIRequest) SetBillTimeEnd(_billTimeEnd string) error {
	r._billTimeEnd = _billTimeEnd
	r.Set("bill_time_end", _billTimeEnd)
	return nil
}

// GetBillTimeEnd BillTimeEnd Getter
func (r TmallServiceSettlementFbBillQueryAPIRequest) GetBillTimeEnd() string {
	return r._billTimeEnd
}

// SetBillTimeStart is BillTimeStart Setter
// 流水开始时间（与截止时间配对使用）
func (r *TmallServiceSettlementFbBillQueryAPIRequest) SetBillTimeStart(_billTimeStart string) error {
	r._billTimeStart = _billTimeStart
	r.Set("bill_time_start", _billTimeStart)
	return nil
}

// GetBillTimeStart BillTimeStart Getter
func (r TmallServiceSettlementFbBillQueryAPIRequest) GetBillTimeStart() string {
	return r._billTimeStart
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TmallServiceSettlementFbBillQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallServiceSettlementFbBillQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 分页页码
func (r *TmallServiceSettlementFbBillQueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TmallServiceSettlementFbBillQueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

var poolTmallServiceSettlementFbBillQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServiceSettlementFbBillQueryRequest()
	},
}

// GetTmallServiceSettlementFbBillQueryRequest 从 sync.Pool 获取 TmallServiceSettlementFbBillQueryAPIRequest
func GetTmallServiceSettlementFbBillQueryAPIRequest() *TmallServiceSettlementFbBillQueryAPIRequest {
	return poolTmallServiceSettlementFbBillQueryAPIRequest.Get().(*TmallServiceSettlementFbBillQueryAPIRequest)
}

// ReleaseTmallServiceSettlementFbBillQueryAPIRequest 将 TmallServiceSettlementFbBillQueryAPIRequest 放入 sync.Pool
func ReleaseTmallServiceSettlementFbBillQueryAPIRequest(v *TmallServiceSettlementFbBillQueryAPIRequest) {
	v.Reset()
	poolTmallServiceSettlementFbBillQueryAPIRequest.Put(v)
}
