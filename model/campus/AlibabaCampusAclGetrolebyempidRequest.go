package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据用户查询角色 API请求
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

// 初始化AlibabaCampusAclGetrolebyempidRequest对象
func NewAlibabaCampusAclGetrolebyempidRequest() *AlibabaCampusAclGetrolebyempidRequest{
    return &AlibabaCampusAclGetrolebyempidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclGetrolebyempidRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.getrolebyempid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclGetrolebyempidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 公司id
func (r *AlibabaCampusAclGetrolebyempidRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclGetrolebyempidRequest) GetCompanyId() int64 {
    return r.companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclGetrolebyempidRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclGetrolebyempidRequest) GetSystemId() string {
    return r.systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclGetrolebyempidRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclGetrolebyempidRequest) GetCampusId() int64 {
    return r.campusId
}
// Param1 Setter
// 用户id
func (r *AlibabaCampusAclGetrolebyempidRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusAclGetrolebyempidRequest) GetParam1() string {
    return r.param1
}
