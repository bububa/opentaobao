package bill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询虚拟账户明细数据(自研发商家专用) APIRequest
taobao.bill.book.bills.get

查询虚拟账户明细数据
*/
type TaobaoBillBookBillsGetRequest struct {
    model.Params

    // 虚拟账户科目编号
    accountId   int64 

    // 明细流水类型:流水类型:101、可用金充值；102、可用金扣除；103、冻结；104、解冻；105、冻结金充值；106、冻结金扣除
    journalTypes   []int64 

    // 记账开始时间
    startTime   string 

    // 记账结束时间,end_time与start_time相差不能超过30天
    endTime   string 

    // 每页大小,建议40~100,不能超过100
    pageSize   int64 

    // 需要返回的字段:bid,account_id,journal_type,amount,book_time,description,gmt_create,gmt_modified ,如果不是以上字段将自动忽略
    fields   string 

    // 页码,传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始
    pageNo   int64 

}

func NewTaobaoBillBookBillsGetRequest() *TaobaoBillBookBillsGetRequest{
    return &TaobaoBillBookBillsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBillBookBillsGetRequest) GetApiMethodName() string {
    return "taobao.bill.book.bills.get"
}

func (r TaobaoBillBookBillsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBillBookBillsGetRequest) SetAccountId(accountId int64) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

func (r TaobaoBillBookBillsGetRequest) GetAccountId() int64 {
    return r.accountId
}

func (r *TaobaoBillBookBillsGetRequest) SetJournalTypes(journalTypes []int64) error {
    r.journalTypes = journalTypes
    r.Set("journal_types", journalTypes)
    return nil
}

func (r TaobaoBillBookBillsGetRequest) GetJournalTypes() []int64 {
    return r.journalTypes
}

func (r *TaobaoBillBookBillsGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoBillBookBillsGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoBillBookBillsGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoBillBookBillsGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoBillBookBillsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoBillBookBillsGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoBillBookBillsGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoBillBookBillsGetRequest) GetFields() string {
    return r.fields
}

func (r *TaobaoBillBookBillsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoBillBookBillsGetRequest) GetPageNo() int64 {
    return r.pageNo
}

