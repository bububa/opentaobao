package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短信查询 API请求
alibaba.aliqin.ta.sms.num.query

查询短信发送揭露
*/
type AlibabaAliqinTaSmsNumQueryRequest struct {
    model.Params
    // 短信发送流水
    _bizId   string
    // 短信接收号码
    _recNum   string
    // 短信发送日期，支持近30天记录查询，格式yyyyMMdd
    _queryDate   string
    // 分页参数,页码
    _currentPage   int64
    // 分页参数，每页数量。最大值50
    _pageSize   int64
}

// 初始化AlibabaAliqinTaSmsNumQueryRequest对象
func NewAlibabaAliqinTaSmsNumQueryRequest() *AlibabaAliqinTaSmsNumQueryRequest{
    return &AlibabaAliqinTaSmsNumQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinTaSmsNumQueryRequest) GetApiMethodName() string {
    return "alibaba.aliqin.ta.sms.num.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinTaSmsNumQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizId Setter
// 短信发送流水
func (r *AlibabaAliqinTaSmsNumQueryRequest) SetBizId(_bizId string) error {
    r._bizId = _bizId
    r.Set("biz_id", _bizId)
    return nil
}

// BizId Getter
func (r AlibabaAliqinTaSmsNumQueryRequest) GetBizId() string {
    return r._bizId
}
// RecNum Setter
// 短信接收号码
func (r *AlibabaAliqinTaSmsNumQueryRequest) SetRecNum(_recNum string) error {
    r._recNum = _recNum
    r.Set("rec_num", _recNum)
    return nil
}

// RecNum Getter
func (r AlibabaAliqinTaSmsNumQueryRequest) GetRecNum() string {
    return r._recNum
}
// QueryDate Setter
// 短信发送日期，支持近30天记录查询，格式yyyyMMdd
func (r *AlibabaAliqinTaSmsNumQueryRequest) SetQueryDate(_queryDate string) error {
    r._queryDate = _queryDate
    r.Set("query_date", _queryDate)
    return nil
}

// QueryDate Getter
func (r AlibabaAliqinTaSmsNumQueryRequest) GetQueryDate() string {
    return r._queryDate
}
// CurrentPage Setter
// 分页参数,页码
func (r *AlibabaAliqinTaSmsNumQueryRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaAliqinTaSmsNumQueryRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// PageSize Setter
// 分页参数，每页数量。最大值50
func (r *AlibabaAliqinTaSmsNumQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAliqinTaSmsNumQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
