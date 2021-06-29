package bill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询虚拟账户明细数据 API请求
taobao.tae.book.bills.get

tae查询虚拟账户明细数据
*/
type TaobaoTaeBookBillsGetRequest struct {
    model.Params
    // 记账开始时间
    _startTime   string
    // 明细流水类型:流水类型:101、可用金充值；102、可用金扣除；103、冻结；104、解冻；105、冻结金充值；106、冻结金扣除
    _journalTypes   []int64
    // 每页大小,建议40~100,不能超过100
    _pageSize   int64
    // 页码,传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始
    _pageNo   int64
    // 虚拟账户科目编号
    _accountId   int64
    // 记账结束时间,end_time与start_time相差不能超过30天
    _endTime   string
    // 需要返回的字段:bid,account_id,journal_type,amount,book_time,description,gmt_create,gmt_modified ,如果不是以上字段将自动忽略
    _fields   []string
}

// 初始化TaobaoTaeBookBillsGetRequest对象
func NewTaobaoTaeBookBillsGetRequest() *TaobaoTaeBookBillsGetRequest{
    return &TaobaoTaeBookBillsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTaeBookBillsGetRequest) GetApiMethodName() string {
    return "taobao.tae.book.bills.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTaeBookBillsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartTime Setter
// 记账开始时间
func (r *TaobaoTaeBookBillsGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoTaeBookBillsGetRequest) GetStartTime() string {
    return r._startTime
}
// JournalTypes Setter
// 明细流水类型:流水类型:101、可用金充值；102、可用金扣除；103、冻结；104、解冻；105、冻结金充值；106、冻结金扣除
func (r *TaobaoTaeBookBillsGetRequest) SetJournalTypes(_journalTypes []int64) error {
    r._journalTypes = _journalTypes
    r.Set("journal_types", _journalTypes)
    return nil
}

// JournalTypes Getter
func (r TaobaoTaeBookBillsGetRequest) GetJournalTypes() []int64 {
    return r._journalTypes
}
// PageSize Setter
// 每页大小,建议40~100,不能超过100
func (r *TaobaoTaeBookBillsGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTaeBookBillsGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 页码,传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始
func (r *TaobaoTaeBookBillsGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTaeBookBillsGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// AccountId Setter
// 虚拟账户科目编号
func (r *TaobaoTaeBookBillsGetRequest) SetAccountId(_accountId int64) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r TaobaoTaeBookBillsGetRequest) GetAccountId() int64 {
    return r._accountId
}
// EndTime Setter
// 记账结束时间,end_time与start_time相差不能超过30天
func (r *TaobaoTaeBookBillsGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoTaeBookBillsGetRequest) GetEndTime() string {
    return r._endTime
}
// Fields Setter
// 需要返回的字段:bid,account_id,journal_type,amount,book_time,description,gmt_create,gmt_modified ,如果不是以上字段将自动忽略
func (r *TaobaoTaeBookBillsGetRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoTaeBookBillsGetRequest) GetFields() []string {
    return r._fields
}
