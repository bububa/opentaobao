package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclqueryallemppermiitemAPIRequest 查询员工全部权限(包括角色下面的权限) API请求
// alibaba.campus.acl.queryallemppermiitem
//
// 查询员工全部权限(包括角色下面的权限)
type AlibabacampusaclqueryallemppermiitemAPIRequest struct {
	model.Params
	// 系统id
	_systemId string
	// 用户账号
	_accountId string
	// 公司id不填默认SYS_000
	_companyId int64
	// 园区id
	_campusId int64
	// 每页多少条
	_page int64
	// 每页记录数
	_pageSize int64
}

// NewAlibabacampusaclqueryallemppermiitemRequest 初始化AlibabacampusaclqueryallemppermiitemAPIRequest对象
func NewAlibabacampusaclqueryallemppermiitemRequest() *AlibabacampusaclqueryallemppermiitemAPIRequest {
	return &AlibabacampusaclqueryallemppermiitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclqueryallemppermiitemAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.queryallemppermiitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclqueryallemppermiitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclqueryallemppermiitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabacampusaclqueryallemppermiitemAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabacampusaclqueryallemppermiitemAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetAccountId is AccountId Setter
// 用户账号
func (r *AlibabacampusaclqueryallemppermiitemAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabacampusaclqueryallemppermiitemAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetCompanyId is CompanyId Setter
// 公司id不填默认SYS_000
func (r *AlibabacampusaclqueryallemppermiitemAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabacampusaclqueryallemppermiitemAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabacampusaclqueryallemppermiitemAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabacampusaclqueryallemppermiitemAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetPage is Page Setter
// 每页多少条
func (r *AlibabacampusaclqueryallemppermiitemAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabacampusaclqueryallemppermiitemAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 每页记录数
func (r *AlibabacampusaclqueryallemppermiitemAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabacampusaclqueryallemppermiitemAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
