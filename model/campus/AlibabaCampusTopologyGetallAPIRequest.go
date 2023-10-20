package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusTopologyGetallAPIRequest 获取管理园区的规则拓扑接口 API请求
// alibaba.campus.topology.getall
//
// 获取所属园区的所有规则拓扑图
type AlibabaCampusTopologyGetallAPIRequest struct {
	model.Params
	// 系统id
	_systemId string
	// 园区id
	_campusId int64
	// 公司id
	_companyId int64
}

// NewAlibabaCampusTopologyGetallRequest 初始化AlibabaCampusTopologyGetallAPIRequest对象
func NewAlibabaCampusTopologyGetallRequest() *AlibabaCampusTopologyGetallAPIRequest {
	return &AlibabaCampusTopologyGetallAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusTopologyGetallAPIRequest) Reset() {
	r._systemId = ""
	r._campusId = 0
	r._companyId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusTopologyGetallAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.topology.getall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusTopologyGetallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusTopologyGetallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemId is SystemId Setter
// 系统id
func (r *AlibabaCampusTopologyGetallAPIRequest) SetSystemId(_systemId string) error {
	r._systemId = _systemId
	r.Set("system_id", _systemId)
	return nil
}

// GetSystemId SystemId Getter
func (r AlibabaCampusTopologyGetallAPIRequest) GetSystemId() string {
	return r._systemId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabaCampusTopologyGetallAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabaCampusTopologyGetallAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetCompanyId is CompanyId Setter
// 公司id
func (r *AlibabaCampusTopologyGetallAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabaCampusTopologyGetallAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

var poolAlibabaCampusTopologyGetallAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusTopologyGetallRequest()
	},
}

// GetAlibabaCampusTopologyGetallRequest 从 sync.Pool 获取 AlibabaCampusTopologyGetallAPIRequest
func GetAlibabaCampusTopologyGetallAPIRequest() *AlibabaCampusTopologyGetallAPIRequest {
	return poolAlibabaCampusTopologyGetallAPIRequest.Get().(*AlibabaCampusTopologyGetallAPIRequest)
}

// ReleaseAlibabaCampusTopologyGetallAPIRequest 将 AlibabaCampusTopologyGetallAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusTopologyGetallAPIRequest(v *AlibabaCampusTopologyGetallAPIRequest) {
	v.Reset()
	poolAlibabaCampusTopologyGetallAPIRequest.Put(v)
}
