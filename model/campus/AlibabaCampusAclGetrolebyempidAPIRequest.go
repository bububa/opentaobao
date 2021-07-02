package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclGetrolebyempidAPIRequest 根据用户查询角色 API请求
// alibaba.campus.acl.getrolebyempid
//
// 根据用户查询角色
type AlibabaCampusAclGetrolebyempidAPIRequest struct {
	model.Params
	// 公司id
	_companyId int64
	// 系统id
	_systemId string
	// 园区id
	_campusId int64
	// 用户id
	_param1 string
}

// NewAlibabaCampusAclGetrolebyempidRequest 初始化AlibabaCampusAclGetrolebyempidAPIRequest对象
func NewAlibabaCampusAclGetrolebyempidRequest() *AlibabaCampusAclGetrolebyempidAPIRequest {
	return &AlibabaCampusAclGetrolebyempidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclGetrolebyempidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.getrolebyempid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclGetrolebyempidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CompanyId Setter
// 公司id
func (r *AlibabaCampusAclGetrolebyempidAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// Get CompanyId Getter
func (r AlibabaCampusAclGetrolebyempidAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// Set is SystemId Setter
// 系统id
func (r *AlibabaCampusAclGetrolebyempidAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// Get SystemId Getter
func (r AlibabaCampusAclGetrolebyempidAPIRequest) GetSystemId() string {
	return r._systemId
}

// Set is CampusId Setter
// 园区id
func (r *AlibabaCampusAclGetrolebyempidAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// Get CampusId Getter
func (r AlibabaCampusAclGetrolebyempidAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// Set is Param1 Setter
// 用户id
func (r *AlibabaCampusAclGetrolebyempidAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r AlibabaCampusAclGetrolebyempidAPIRequest) GetParam1() string {
	return r._param1
}
