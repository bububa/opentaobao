package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclCheckemproleAPIRequest 校验用户是否有该角色 API请求
// alibaba.campus.acl.checkemprole
//
// 校验用户是否有该权限
type AlibabaCampusAclCheckemproleAPIRequest struct {
	model.Params
	// 系统id
	_systemId string
	// 员工账号
	_accountId string
	// 角色id
	_itemKey string
	// 公司id不填默认为SYS_000
	_companyId int64
	// 园区id
	_campusId int64
}

// NewAlibabaCampusAclCheckemproleRequest 初始化AlibabaCampusAclCheckemproleAPIRequest对象
func NewAlibabaCampusAclCheckemproleRequest() *AlibabaCampusAclCheckemproleAPIRequest {
	return &AlibabaCampusAclCheckemproleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclCheckemproleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.checkemprole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclCheckemproleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclCheckemproleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabaCampusAclCheckemproleAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabaCampusAclCheckemproleAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetAccountId is AccountId Setter
// 员工账号
func (r *AlibabaCampusAclCheckemproleAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabaCampusAclCheckemproleAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetItemKey is ItemKey Setter
// 角色id
func (r *AlibabaCampusAclCheckemproleAPIRequest) SetItemKey(_itemKey string) error {
	r._itemKey = _itemKey
	r.Set("item_key", _itemKey)
	return nil
}

// GetItemKey ItemKey Getter
func (r AlibabaCampusAclCheckemproleAPIRequest) GetItemKey() string {
	return r._itemKey
}

// SetCompanyId is CompanyId Setter
// 公司id不填默认为SYS_000
func (r *AlibabaCampusAclCheckemproleAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabaCampusAclCheckemproleAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabaCampusAclCheckemproleAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabaCampusAclCheckemproleAPIRequest) GetCampusId() int64 {
	return r._campusId
}
