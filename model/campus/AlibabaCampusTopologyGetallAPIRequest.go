package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampustopologygetallAPIRequest 获取管理园区的规则拓扑接口 API请求
// alibaba.campus.topology.getall
//
// 获取所属园区的所有规则拓扑图
type AlibabacampustopologygetallAPIRequest struct {
	model.Params
	// 系统id
	_systemId string
	// 园区id
	_campusId int64
	// 公司id
	_companyId int64
}

// NewAlibabacampustopologygetallRequest 初始化AlibabacampustopologygetallAPIRequest对象
func NewAlibabacampustopologygetallRequest() *AlibabacampustopologygetallAPIRequest {
	return &AlibabacampustopologygetallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampustopologygetallAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.topology.getall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampustopologygetallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampustopologygetallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabacampustopologygetallAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabacampustopologygetallAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabacampustopologygetallAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabacampustopologygetallAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetCompanyId is CompanyId Setter
// 公司id
func (r *AlibabacampustopologygetallAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabacampustopologygetallAPIRequest) GetCompanyId() int64 {
	return r._companyId
}
