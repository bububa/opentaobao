package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据用户查询角色 APIRequest
alibaba.campus.acl.getrolebyempid

根据用户查询角色
*/
type AlibabaCampusAclGetrolebyempidRequest struct {
    model.Params

    // 公司id
    companyId   int64 

    // 系统id
    systemId   string 

    // 园区id
    campusId   int64 

    // 用户id
    param1   string 

}

func NewAlibabaCampusAclGetrolebyempidRequest() *AlibabaCampusAclGetrolebyempidRequest{
    return &AlibabaCampusAclGetrolebyempidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclGetrolebyempidRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.getrolebyempid"
}

func (r AlibabaCampusAclGetrolebyempidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclGetrolebyempidRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaCampusAclGetrolebyempidRequest) GetCompanyId() int64 {
    return r.companyId
}

func (r *AlibabaCampusAclGetrolebyempidRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

func (r AlibabaCampusAclGetrolebyempidRequest) GetSystemId() string {
    return r.systemId
}

func (r *AlibabaCampusAclGetrolebyempidRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaCampusAclGetrolebyempidRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaCampusAclGetrolebyempidRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaCampusAclGetrolebyempidRequest) GetParam1() string {
    return r.param1
}

