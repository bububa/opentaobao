package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户的菜单 APIRequest
alibaba.campus.acl.getmenubyempid

查询用户的菜单
*/
type AlibabaCampusAclGetmenubyempidRequest struct {
    model.Params

    // 账户id
    userId   int64 

    // 系统id
    systemId   string 

    // 公司id
    companyId   int64 

    // 园区id
    campusId   int64 

}

func NewAlibabaCampusAclGetmenubyempidRequest() *AlibabaCampusAclGetmenubyempidRequest{
    return &AlibabaCampusAclGetmenubyempidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclGetmenubyempidRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.getmenubyempid"
}

func (r AlibabaCampusAclGetmenubyempidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclGetmenubyempidRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaCampusAclGetmenubyempidRequest) GetUserId() int64 {
    return r.userId
}

func (r *AlibabaCampusAclGetmenubyempidRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

func (r AlibabaCampusAclGetmenubyempidRequest) GetSystemId() string {
    return r.systemId
}

func (r *AlibabaCampusAclGetmenubyempidRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaCampusAclGetmenubyempidRequest) GetCompanyId() int64 {
    return r.companyId
}

func (r *AlibabaCampusAclGetmenubyempidRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaCampusAclGetmenubyempidRequest) GetCampusId() int64 {
    return r.campusId
}

