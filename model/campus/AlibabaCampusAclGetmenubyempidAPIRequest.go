package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclGetmenubyempidAPIRequest 查询用户的菜单 API请求
// alibaba.campus.acl.getmenubyempid
//
// 查询用户的菜单
type AlibabaCampusAclGetmenubyempidAPIRequest struct {
	model.Params
	// 账户id
	_userId int64
	// 系统id
	_systemId string
	// 公司id
	_companyId int64
	// 园区id
	_campusId int64
}

// NewAlibabaCampusAclGetmenubyempidRequest 初始化AlibabaCampusAclGetmenubyempidAPIRequest对象
func NewAlibabaCampusAclGetmenubyempidRequest() *AlibabaCampusAclGetmenubyempidAPIRequest {
	return &AlibabaCampusAclGetmenubyempidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclGetmenubyempidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.getmenubyempid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclGetmenubyempidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUserId is UserId Setter
// 账户id
func (r *AlibabaCampusAclGetmenubyempidAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaCampusAclGetmenubyempidAPIRequest) GetUserId() int64 {
	return r._userId
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabaCampusAclGetmenubyempidAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabaCampusAclGetmenubyempidAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetCompanyId is CompanyId Setter
// 公司id
func (r *AlibabaCampusAclGetmenubyempidAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabaCampusAclGetmenubyempidAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabaCampusAclGetmenubyempidAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabaCampusAclGetmenubyempidAPIRequest) GetCampusId() int64 {
	return r._campusId
}
