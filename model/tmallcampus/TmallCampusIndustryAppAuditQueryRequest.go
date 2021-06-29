package tmallcampus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫校园查询学生认证信息 API请求
tmall.campus.industry.app.audit.query

天猫校园查询学生认证信息
*/
type TmallCampusIndustryAppAuditQueryRequest struct {
    model.Params
    // 调用来源
    _source   string
}

// 初始化TmallCampusIndustryAppAuditQueryRequest对象
func NewTmallCampusIndustryAppAuditQueryRequest() *TmallCampusIndustryAppAuditQueryRequest{
    return &TmallCampusIndustryAppAuditQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCampusIndustryAppAuditQueryRequest) GetApiMethodName() string {
    return "tmall.campus.industry.app.audit.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallCampusIndustryAppAuditQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Source Setter
// 调用来源
func (r *TmallCampusIndustryAppAuditQueryRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TmallCampusIndustryAppAuditQueryRequest) GetSource() string {
    return r._source
}
