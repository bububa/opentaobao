package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短信发送记录查询 APIRequest
alibaba.aliqin.fc.sms.num.query

短信发送记录查询。
*/
type AlibabaAliqinFcSmsNumQueryRequest struct {
    model.Params

    // 短信发送流水
    bizId   string 

    // 短信接收号码
    recNum   string 

    // 短信发送日期，支持近30天记录查询，格式yyyyMMdd
    queryDate   string 

    // 分页参数,页码
    currentPage   int64 

    // 分页参数，每页数量。最大值50
    pageSize   int64 

}

func NewAlibabaAliqinFcSmsNumQueryRequest() *AlibabaAliqinFcSmsNumQueryRequest{
    return &AlibabaAliqinFcSmsNumQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFcSmsNumQueryRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.sms.num.query"
}

func (r AlibabaAliqinFcSmsNumQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFcSmsNumQueryRequest) SetBizId(bizId string) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

func (r AlibabaAliqinFcSmsNumQueryRequest) GetBizId() string {
    return r.bizId
}

func (r *AlibabaAliqinFcSmsNumQueryRequest) SetRecNum(recNum string) error {
    r.recNum = recNum
    r.Set("rec_num", recNum)
    return nil
}

func (r AlibabaAliqinFcSmsNumQueryRequest) GetRecNum() string {
    return r.recNum
}

func (r *AlibabaAliqinFcSmsNumQueryRequest) SetQueryDate(queryDate string) error {
    r.queryDate = queryDate
    r.Set("query_date", queryDate)
    return nil
}

func (r AlibabaAliqinFcSmsNumQueryRequest) GetQueryDate() string {
    return r.queryDate
}

func (r *AlibabaAliqinFcSmsNumQueryRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlibabaAliqinFcSmsNumQueryRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *AlibabaAliqinFcSmsNumQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAliqinFcSmsNumQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

