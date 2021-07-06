package bill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaeBookBillsGetAPIRequest tae查询虚拟账户明细数据 API请求
// taobao.tae.book.bills.get
//
// tae查询虚拟账户明细数据
type TaobaoTaeBookBillsGetAPIRequest struct {
	model.Params
	// 明细流水类型:流水类型:101、可用金充值；102、可用金扣除；103、冻结；104、解冻；105、冻结金充值；106、冻结金扣除
	_journalTypes []int64
	// 需要返回的字段:bid,account_id,journal_type,amount,book_time,description,gmt_create,gmt_modified ,如果不是以上字段将自动忽略
	_fields []string
	// 记账开始时间
	_startTime string
	// 记账结束时间,end_time与start_time相差不能超过30天
	_endTime string
	// 每页大小,建议40~100,不能超过100
	_pageSize int64
	// 页码,传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始
	_pageNo int64
	// 虚拟账户科目编号
	_accountId int64
}

// NewTaobaoTaeBookBillsGetRequest 初始化TaobaoTaeBookBillsGetAPIRequest对象
func NewTaobaoTaeBookBillsGetRequest() *TaobaoTaeBookBillsGetAPIRequest {
	return &TaobaoTaeBookBillsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaeBookBillsGetAPIRequest) GetApiMethodName() string {
	return "taobao.tae.book.bills.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaeBookBillsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetJournalTypes is JournalTypes Setter
// 明细流水类型:流水类型:101、可用金充值；102、可用金扣除；103、冻结；104、解冻；105、冻结金充值；106、冻结金扣除
func (r *TaobaoTaeBookBillsGetAPIRequest) SetJournalTypes(_journalTypes []int64) error {
	r._journalTypes = _journalTypes
	r.Set("journal_types", _journalTypes)
	return nil
}

// GetJournalTypes JournalTypes Getter
func (r TaobaoTaeBookBillsGetAPIRequest) GetJournalTypes() []int64 {
	return r._journalTypes
}

// SetFields is Fields Setter
// 需要返回的字段:bid,account_id,journal_type,amount,book_time,description,gmt_create,gmt_modified ,如果不是以上字段将自动忽略
func (r *TaobaoTaeBookBillsGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTaeBookBillsGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetStartTime is StartTime Setter
// 记账开始时间
func (r *TaobaoTaeBookBillsGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoTaeBookBillsGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 记账结束时间,end_time与start_time相差不能超过30天
func (r *TaobaoTaeBookBillsGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoTaeBookBillsGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPageSize is PageSize Setter
// 每页大小,建议40~100,不能超过100
func (r *TaobaoTaeBookBillsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTaeBookBillsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码,传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始
func (r *TaobaoTaeBookBillsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoTaeBookBillsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetAccountId is AccountId Setter
// 虚拟账户科目编号
func (r *TaobaoTaeBookBillsGetAPIRequest) SetAccountId(_accountId int64) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r TaobaoTaeBookBillsGetAPIRequest) GetAccountId() int64 {
	return r._accountId
}
