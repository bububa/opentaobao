package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettlementFbBillDetailQueryAPIRequest 服务商工单结算对账查询-流水查询 API请求
// tmall.service.settlement.fb.bill.detail.query
//
// 服务商工单结算对账查询-流水查询，用于查询服务工单费用流水，含服务费、退款、分成、提现等。
type TmallServiceSettlementFbBillDetailQueryAPIRequest struct {
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

// NewTmallServiceSettlementFbBillDetailQueryRequest 初始化TmallServiceSettlementFbBillDetailQueryAPIRequest对象
func NewTmallServiceSettlementFbBillDetailQueryRequest() *TmallServiceSettlementFbBillDetailQueryAPIRequest {
	return &TmallServiceSettlementFbBillDetailQueryAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServiceSettlementFbBillDetailQueryAPIRequest) Reset() {
	r._payTimeStart = ""
	r._payTimeEnd = ""
	r._billTimeEnd = ""
	r._billTimeStart = ""
	r._pageSize = 0
	r._pageNum = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServiceSettlementFbBillDetailQueryAPIRequest) GetApiMethodName() string {
	return "tmall.service.settlement.fb.bill.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServiceSettlementFbBillDetailQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServiceSettlementFbBillDetailQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayTimeStart is PayTimeStart Setter
// 提现开始时间（如果设置了开始时间，需要同时设置截止时间）
func (r *TmallServiceSettlementFbBillDetailQueryAPIRequest) SetPayTimeStart(_payTimeStart string) error {
	r._payTimeStart = _payTimeStart
	r.Set("pay_time_start", _payTimeStart)
	return nil
}

// GetPayTimeStart PayTimeStart Getter
func (r TmallServiceSettlementFbBillDetailQueryAPIRequest) GetPayTimeStart() string {
	return r._payTimeStart
}

// SetPayTimeEnd is PayTimeEnd Setter
// 提现截止时间
func (r *TmallServiceSettlementFbBillDetailQueryAPIRequest) SetPayTimeEnd(_payTimeEnd string) error {
	r._payTimeEnd = _payTimeEnd
	r.Set("pay_time_end", _payTimeEnd)
	return nil
}

// GetPayTimeEnd PayTimeEnd Getter
func (r TmallServiceSettlementFbBillDetailQueryAPIRequest) GetPayTimeEnd() string {
	return r._payTimeEnd
}

// SetBillTimeEnd is BillTimeEnd Setter
// 流水截止时间
func (r *TmallServiceSettlementFbBillDetailQueryAPIRequest) SetBillTimeEnd(_billTimeEnd string) error {
	r._billTimeEnd = _billTimeEnd
	r.Set("bill_time_end", _billTimeEnd)
	return nil
}

// GetBillTimeEnd BillTimeEnd Getter
func (r TmallServiceSettlementFbBillDetailQueryAPIRequest) GetBillTimeEnd() string {
	return r._billTimeEnd
}

// SetBillTimeStart is BillTimeStart Setter
// 流水开始时间（如果设置了开始时间，需要同时设置截止时间）
func (r *TmallServiceSettlementFbBillDetailQueryAPIRequest) SetBillTimeStart(_billTimeStart string) error {
	r._billTimeStart = _billTimeStart
	r.Set("bill_time_start", _billTimeStart)
	return nil
}

// GetBillTimeStart BillTimeStart Getter
func (r TmallServiceSettlementFbBillDetailQueryAPIRequest) GetBillTimeStart() string {
	return r._billTimeStart
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TmallServiceSettlementFbBillDetailQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallServiceSettlementFbBillDetailQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 分页页码
func (r *TmallServiceSettlementFbBillDetailQueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TmallServiceSettlementFbBillDetailQueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

var poolTmallServiceSettlementFbBillDetailQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServiceSettlementFbBillDetailQueryRequest()
	},
}

// GetTmallServiceSettlementFbBillDetailQueryRequest 从 sync.Pool 获取 TmallServiceSettlementFbBillDetailQueryAPIRequest
func GetTmallServiceSettlementFbBillDetailQueryAPIRequest() *TmallServiceSettlementFbBillDetailQueryAPIRequest {
	return poolTmallServiceSettlementFbBillDetailQueryAPIRequest.Get().(*TmallServiceSettlementFbBillDetailQueryAPIRequest)
}

// ReleaseTmallServiceSettlementFbBillDetailQueryAPIRequest 将 TmallServiceSettlementFbBillDetailQueryAPIRequest 放入 sync.Pool
func ReleaseTmallServiceSettlementFbBillDetailQueryAPIRequest(v *TmallServiceSettlementFbBillDetailQueryAPIRequest) {
	v.Reset()
	poolTmallServiceSettlementFbBillDetailQueryAPIRequest.Put(v)
}
