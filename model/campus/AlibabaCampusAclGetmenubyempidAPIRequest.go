package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclgetmenubyempidAPIRequest 查询用户的菜单 API请求
// alibaba.campus.acl.getmenubyempid
//
// 查询用户的菜单
type AlibabacampusaclgetmenubyempidAPIRequest struct {
	model.Params
	// 系统id
	_systemId string
	// 账户id
	_userId int64
	// 公司id
	_companyId int64
	// 园区id
	_campusId int64
}

// NewAlibabacampusaclgetmenubyempidRequest 初始化AlibabacampusaclgetmenubyempidAPIRequest对象
func NewAlibabacampusaclgetmenubyempidRequest() *AlibabacampusaclgetmenubyempidAPIRequest {
	return &AlibabacampusaclgetmenubyempidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclgetmenubyempidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.getmenubyempid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclgetmenubyempidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclgetmenubyempidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabacampusaclgetmenubyempidAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabacampusaclgetmenubyempidAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetUserId is UserId Setter
// 账户id
func (r *AlibabacampusaclgetmenubyempidAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabacampusaclgetmenubyempidAPIRequest) GetUserId() int64 {
	return r._userId
}

// SetCompanyId is CompanyId Setter
// 公司id
func (r *AlibabacampusaclgetmenubyempidAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabacampusaclgetmenubyempidAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabacampusaclgetmenubyempidAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabacampusaclgetmenubyempidAPIRequest) GetCampusId() int64 {
	return r._campusId
}
