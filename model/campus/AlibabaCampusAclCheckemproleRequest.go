package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否有该角色 API请求
alibaba.campus.acl.checkemprole

校验用户是否有该权限
*/
type AlibabaCampusAclCheckemproleRequest struct {
    model.Params
    // 公司id不填默认为SYS_000
    _companyId   int64
    // 系统id
    _systemId   string
    // 园区id
    _campusId   int64
    // 员工账号
    _accountId   string
    // 角色id
    _itemKey   string
}

// 初始化AlibabaCampusAclCheckemproleRequest对象
func NewAlibabaCampusAclCheckemproleRequest() *AlibabaCampusAclCheckemproleRequest{
    return &AlibabaCampusAclCheckemproleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclCheckemproleRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.checkemprole"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclCheckemproleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 公司id不填默认为SYS_000
func (r *AlibabaCampusAclCheckemproleRequest) SetCompanyId(_companyId int64) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclCheckemproleRequest) GetCompanyId() int64 {
    return r._companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclCheckemproleRequest) SetSystemId(_systemId string) error {
    r._systemId = _systemId
    r.Set("system_id", _systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclCheckemproleRequest) GetSystemId() string {
    return r._systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclCheckemproleRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclCheckemproleRequest) GetCampusId() int64 {
    return r._campusId
}
// AccountId Setter
// 员工账号
func (r *AlibabaCampusAclCheckemproleRequest) SetAccountId(_accountId string) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r AlibabaCampusAclCheckemproleRequest) GetAccountId() string {
    return r._accountId
}
// ItemKey Setter
// 角色id
func (r *AlibabaCampusAclCheckemproleRequest) SetItemKey(_itemKey string) error {
    r._itemKey = _itemKey
    r.Set("item_key", _itemKey)
    return nil
}

// ItemKey Getter
func (r AlibabaCampusAclCheckemproleRequest) GetItemKey() string {
    return r._itemKey
}
