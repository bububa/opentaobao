package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusTopologyGetallAPIRequest 获取管理园区的规则拓扑接口 API请求
// alibaba.campus.topology.getall
//
// 获取所属园区的所有规则拓扑图
type AlibabaCampusTopologyGetallAPIRequest struct {
	model.Params
	// 园区id
	_campusId int64
	// 公司id
	_companyId int64
	// 系统id
	_systemId string
}

// NewAlibabaCampusTopologyGetallRequest 初始化AlibabaCampusTopologyGetallAPIRequest对象
func NewAlibabaCampusTopologyGetallRequest() *AlibabaCampusTopologyGetallAPIRequest {
	return &AlibabaCampusTopologyGetallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusTopologyGetallAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.topology.getall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusTopologyGetallAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CampusId Setter
// 园区id
func (r *AlibabaCampusTopologyGetallAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// Get CampusId Getter
func (r AlibabaCampusTopologyGetallAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// Set is CompanyId Setter
// 公司id
func (r *AlibabaCampusTopologyGetallAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// Get CompanyId Getter
func (r AlibabaCampusTopologyGetallAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// Set is SystemId Setter
// 系统id
func (r *AlibabaCampusTopologyGetallAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// Get SystemId Getter
func (r AlibabaCampusTopologyGetallAPIRequest) GetSystemId() string {
	return r._systemId
}
