package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户的菜单 API请求
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

// 初始化AlibabaCampusAclGetmenubyempidRequest对象
func NewAlibabaCampusAclGetmenubyempidRequest() *AlibabaCampusAclGetmenubyempidRequest{
    return &AlibabaCampusAclGetmenubyempidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclGetmenubyempidRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.getmenubyempid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclGetmenubyempidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 账户id
func (r *AlibabaCampusAclGetmenubyempidRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclGetmenubyempidRequest) GetUserId() int64 {
    return r.userId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclGetmenubyempidRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclGetmenubyempidRequest) GetSystemId() string {
    return r.systemId
}
// CompanyId Setter
// 公司id
func (r *AlibabaCampusAclGetmenubyempidRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclGetmenubyempidRequest) GetCompanyId() int64 {
    return r.companyId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclGetmenubyempidRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclGetmenubyempidRequest) GetCampusId() int64 {
    return r.campusId
}
