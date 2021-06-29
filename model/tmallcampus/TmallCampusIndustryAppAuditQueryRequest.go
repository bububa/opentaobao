package tmallcampus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫校园查询学生认证信息 APIRequest
tmall.campus.industry.app.audit.query

天猫校园查询学生认证信息
*/
type TmallCampusIndustryAppAuditQueryRequest struct {
    model.Params

    // 调用来源
    source   string 

}

func NewTmallCampusIndustryAppAuditQueryRequest() *TmallCampusIndustryAppAuditQueryRequest{
    return &TmallCampusIndustryAppAuditQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCampusIndustryAppAuditQueryRequest) GetApiMethodName() string {
    return "tmall.campus.industry.app.audit.query"
}

func (r TmallCampusIndustryAppAuditQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCampusIndustryAppAuditQueryRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r TmallCampusIndustryAppAuditQueryRequest) GetSource() string {
    return r.source
}

