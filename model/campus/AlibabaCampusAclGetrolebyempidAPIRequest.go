package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclgetrolebyempidAPIRequest 根据用户查询角色 API请求
// alibaba.campus.acl.getrolebyempid
//
// 根据用户查询角色
type AlibabacampusaclgetrolebyempidAPIRequest struct {
	model.Params
	// 系统id
	_systemId string
	// 用户id
	_param1 string
	// 公司id
	_companyId int64
	// 园区id
	_campusId int64
}

// NewAlibabacampusaclgetrolebyempidRequest 初始化AlibabacampusaclgetrolebyempidAPIRequest对象
func NewAlibabacampusaclgetrolebyempidRequest() *AlibabacampusaclgetrolebyempidAPIRequest {
	return &AlibabacampusaclgetrolebyempidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclgetrolebyempidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.getrolebyempid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclgetrolebyempidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclgetrolebyempidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabacampusaclgetrolebyempidAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabacampusaclgetrolebyempidAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetParam1 is Param1 Setter
// 用户id
func (r *AlibabacampusaclgetrolebyempidAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusaclgetrolebyempidAPIRequest) GetParam1() string {
	return r._param1
}

// SetCompanyId is CompanyId Setter
// 公司id
func (r *AlibabacampusaclgetrolebyempidAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabacampusaclgetrolebyempidAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabacampusaclgetrolebyempidAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabacampusaclgetrolebyempidAPIRequest) GetCampusId() int64 {
	return r._campusId
}
