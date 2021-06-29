package bill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询虚拟账户明细数据(自研发商家专用) API请求
taobao.bill.book.bills.get

查询虚拟账户明细数据
*/
type TaobaoBillBookBillsGetRequest struct {
    model.Params
    // 虚拟账户科目编号
    _accountId   int64
    // 明细流水类型:流水类型:101、可用金充值；102、可用金扣除；103、冻结；104、解冻；105、冻结金充值；106、冻结金扣除
    _journalTypes   []int64
    // 记账开始时间
    _startTime   string
    // 记账结束时间,end_time与start_time相差不能超过30天
    _endTime   string
    // 每页大小,建议40~100,不能超过100
    _pageSize   int64
    // 需要返回的字段:bid,account_id,journal_type,amount,book_time,description,gmt_create,gmt_modified ,如果不是以上字段将自动忽略
    _fields   string
    // 页码,传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始
    _pageNo   int64
}

// 初始化TaobaoBillBookBillsGetRequest对象
func NewTaobaoBillBookBillsGetRequest() *TaobaoBillBookBillsGetRequest{
    return &TaobaoBillBookBillsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBillBookBillsGetRequest) GetApiMethodName() string {
    return "taobao.bill.book.bills.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBillBookBillsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountId Setter
// 虚拟账户科目编号
func (r *TaobaoBillBookBillsGetRequest) SetAccountId(_accountId int64) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r TaobaoBillBookBillsGetRequest) GetAccountId() int64 {
    return r._accountId
}
// JournalTypes Setter
// 明细流水类型:流水类型:101、可用金充值；102、可用金扣除；103、冻结；104、解冻；105、冻结金充值；106、冻结金扣除
func (r *TaobaoBillBookBillsGetRequest) SetJournalTypes(_journalTypes []int64) error {
    r._journalTypes = _journalTypes
    r.Set("journal_types", _journalTypes)
    return nil
}

// JournalTypes Getter
func (r TaobaoBillBookBillsGetRequest) GetJournalTypes() []int64 {
    return r._journalTypes
}
// StartTime Setter
// 记账开始时间
func (r *TaobaoBillBookBillsGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoBillBookBillsGetRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 记账结束时间,end_time与start_time相差不能超过30天
func (r *TaobaoBillBookBillsGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoBillBookBillsGetRequest) GetEndTime() string {
    return r._endTime
}
// PageSize Setter
// 每页大小,建议40~100,不能超过100
func (r *TaobaoBillBookBillsGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoBillBookBillsGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// Fields Setter
// 需要返回的字段:bid,account_id,journal_type,amount,book_time,description,gmt_create,gmt_modified ,如果不是以上字段将自动忽略
func (r *TaobaoBillBookBillsGetRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoBillBookBillsGetRequest) GetFields() string {
    return r._fields
}
// PageNo Setter
// 页码,传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始
func (r *TaobaoBillBookBillsGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoBillBookBillsGetRequest) GetPageNo() int64 {
    return r._pageNo
}
