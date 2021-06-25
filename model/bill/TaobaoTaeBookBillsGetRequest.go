package bill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询虚拟账户明细数据 APIRequest
taobao.tae.book.bills.get

tae查询虚拟账户明细数据
*/
type TaobaoTaeBookBillsGetRequest struct {
    model.Params

    // 记账开始时间
    startTime   string 

    // 明细流水类型:流水类型:101、可用金充值；102、可用金扣除；103、冻结；104、解冻；105、冻结金充值；106、冻结金扣除
    journalTypes   []Number 

    // 每页大小,建议40~100,不能超过100
    pageSize   int64 

    // 页码,传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始
    pageNo   int64 

    // 虚拟账户科目编号
    accountId   int64 

    // 记账结束时间,end_time与start_time相差不能超过30天
    endTime   string 

    // 需要返回的字段:bid,account_id,journal_type,amount,book_time,description,gmt_create,gmt_modified ,如果不是以上字段将自动忽略
    fields   []String 

}

func NewTaobaoTaeBookBillsGetRequest() *TaobaoTaeBookBillsGetRequest{
    return &TaobaoTaeBookBillsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTaeBookBillsGetRequest) GetApiMethodName() string {
    return "taobao.tae.book.bills.get"
}

func (r TaobaoTaeBookBillsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTaeBookBillsGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoTaeBookBillsGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoTaeBookBillsGetRequest) SetJournalTypes(journalTypes []Number) error {
    r.journalTypes = journalTypes
    r.Set("journal_types", journalTypes)
    return nil
}

func (r TaobaoTaeBookBillsGetRequest) GetJournalTypes() []Number {
    return r.journalTypes
}

func (r *TaobaoTaeBookBillsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoTaeBookBillsGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoTaeBookBillsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoTaeBookBillsGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoTaeBookBillsGetRequest) SetAccountId(accountId int64) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

func (r TaobaoTaeBookBillsGetRequest) GetAccountId() int64 {
    return r.accountId
}

func (r *TaobaoTaeBookBillsGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoTaeBookBillsGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoTaeBookBillsGetRequest) SetFields(fields []String) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoTaeBookBillsGetRequest) GetFields() []String {
    return r.fields
}

