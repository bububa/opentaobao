package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短信查询 APIRequest
alibaba.aliqin.ta.sms.num.query

查询短信发送揭露
*/
type AlibabaAliqinTaSmsNumQueryRequest struct {
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

func NewAlibabaAliqinTaSmsNumQueryRequest() *AlibabaAliqinTaSmsNumQueryRequest{
    return &AlibabaAliqinTaSmsNumQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinTaSmsNumQueryRequest) GetApiMethodName() string {
    return "alibaba.aliqin.ta.sms.num.query"
}

func (r AlibabaAliqinTaSmsNumQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinTaSmsNumQueryRequest) SetBizId(bizId string) error {
    r.bizId = bizId
    r.Set("biz_id", bizId)
    return nil
}

func (r AlibabaAliqinTaSmsNumQueryRequest) GetBizId() string {
    return r.bizId
}

func (r *AlibabaAliqinTaSmsNumQueryRequest) SetRecNum(recNum string) error {
    r.recNum = recNum
    r.Set("rec_num", recNum)
    return nil
}

func (r AlibabaAliqinTaSmsNumQueryRequest) GetRecNum() string {
    return r.recNum
}

func (r *AlibabaAliqinTaSmsNumQueryRequest) SetQueryDate(queryDate string) error {
    r.queryDate = queryDate
    r.Set("query_date", queryDate)
    return nil
}

func (r AlibabaAliqinTaSmsNumQueryRequest) GetQueryDate() string {
    return r.queryDate
}

func (r *AlibabaAliqinTaSmsNumQueryRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlibabaAliqinTaSmsNumQueryRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *AlibabaAliqinTaSmsNumQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAliqinTaSmsNumQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

