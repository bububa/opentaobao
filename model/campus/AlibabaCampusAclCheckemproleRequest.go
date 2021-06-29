package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否有该角色 APIRequest
alibaba.campus.acl.checkemprole

校验用户是否有该权限
*/
type AlibabaCampusAclCheckemproleRequest struct {
    model.Params

    // 公司id不填默认为SYS_000
    companyId   int64 

    // 系统id
    systemId   string 

    // 园区id
    campusId   int64 

    // 员工账号
    accountId   string 

    // 角色id
    itemKey   string 

}

func NewAlibabaCampusAclCheckemproleRequest() *AlibabaCampusAclCheckemproleRequest{
    return &AlibabaCampusAclCheckemproleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclCheckemproleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.checkemprole"
}

func (r AlibabaCampusAclCheckemproleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclCheckemproleRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaCampusAclCheckemproleRequest) GetCompanyId() int64 {
    return r.companyId
}

func (r *AlibabaCampusAclCheckemproleRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

func (r AlibabaCampusAclCheckemproleRequest) GetSystemId() string {
    return r.systemId
}

func (r *AlibabaCampusAclCheckemproleRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaCampusAclCheckemproleRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaCampusAclCheckemproleRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

func (r AlibabaCampusAclCheckemproleRequest) GetAccountId() string {
    return r.accountId
}

func (r *AlibabaCampusAclCheckemproleRequest) SetItemKey(itemKey string) error {
    r.itemKey = itemKey
    r.Set("item_key", itemKey)
    return nil
}

func (r AlibabaCampusAclCheckemproleRequest) GetItemKey() string {
    return r.itemKey
}

