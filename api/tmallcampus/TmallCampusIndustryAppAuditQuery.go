package tmallcampus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallcampus"
)

/* 
天猫校园查询学生认证信息 
tmall.campus.industry.app.audit.query

天猫校园查询学生认证信息
*/
func TmallCampusIndustryAppAuditQuery(clt *core.SDKClient, req *tmallcampus.TmallCampusIndustryAppAuditQueryAPIRequest, session string) (*tmallcampus.TmallCampusIndustryAppAuditQueryAPIResponse, error) {
    var resp tmallcampus.TmallCampusIndustryAppAuditQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
