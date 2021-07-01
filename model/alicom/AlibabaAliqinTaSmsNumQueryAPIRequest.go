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
type AlibabaAliqinTaSmsNumQueryAPIRequest struct {
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

// 初始化AlibabaAliqinTaSmsNumQueryAPIRequest对象
func NewAlibabaAliqinTaSmsNumQueryRequest() *AlibabaAliqinTaSmsNumQueryAPIRequest{
    return &AlibabaAliqinTaSmsNumQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.aliqin.ta.sms.num.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizId Setter
// 短信发送流水
func (r *AlibabaAliqinTaSmsNumQueryAPIRequest) SetBizId(_bizId string) error {
    r._bizId = _bizId
    r.Set("biz_id", _bizId)
    return nil
}

// BizId Getter
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetBizId() string {
    return r._bizId
}
// RecNum Setter
// 短信接收号码
func (r *AlibabaAliqinTaSmsNumQueryAPIRequest) SetRecNum(_recNum string) error {
    r._recNum = _recNum
    r.Set("rec_num", _recNum)
    return nil
}

// RecNum Getter
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetRecNum() string {
    return r._recNum
}
// QueryDate Setter
// 短信发送日期，支持近30天记录查询，格式yyyyMMdd
func (r *AlibabaAliqinTaSmsNumQueryAPIRequest) SetQueryDate(_queryDate string) error {
    r._queryDate = _queryDate
    r.Set("query_date", _queryDate)
    return nil
}

// QueryDate Getter
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetQueryDate() string {
    return r._queryDate
}
// CurrentPage Setter
// 分页参数,页码
func (r *AlibabaAliqinTaSmsNumQueryAPIRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// PageSize Setter
// 分页参数，每页数量。最大值50
func (r *AlibabaAliqinTaSmsNumQueryAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAliqinTaSmsNumQueryAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
