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
    _userId   int64
    // 系统id
    _systemId   string
    // 公司id
    _companyId   int64
    // 园区id
    _campusId   int64
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
func (r *AlibabaCampusAclGetmenubyempidRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclGetmenubyempidRequest) GetUserId() int64 {
    return r._userId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclGetmenubyempidRequest) SetSystemId(_systemId string) error {
    r._systemId = _systemId
    r.Set("system_id", _systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclGetmenubyempidRequest) GetSystemId() string {
    return r._systemId
}
// CompanyId Setter
// 公司id
func (r *AlibabaCampusAclGetmenubyempidRequest) SetCompanyId(_companyId int64) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclGetmenubyempidRequest) GetCompanyId() int64 {
    return r._companyId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclGetmenubyempidRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclGetmenubyempidRequest) GetCampusId() int64 {
    return r._campusId
}
