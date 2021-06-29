package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取管理园区的规则拓扑接口 API请求
alibaba.campus.topology.getall

获取所属园区的所有规则拓扑图
*/
type AlibabaCampusTopologyGetallRequest struct {
    model.Params
    // 园区id
    _campusId   int64
    // 公司id
    _companyId   int64
    // 系统id
    _systemId   string
}

// 初始化AlibabaCampusTopologyGetallRequest对象
func NewAlibabaCampusTopologyGetallRequest() *AlibabaCampusTopologyGetallRequest{
    return &AlibabaCampusTopologyGetallRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusTopologyGetallRequest) GetApiMethodName() string {
    return "alibaba.campus.topology.getall"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusTopologyGetallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusTopologyGetallRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusTopologyGetallRequest) GetCampusId() int64 {
    return r._campusId
}
// CompanyId Setter
// 公司id
func (r *AlibabaCampusTopologyGetallRequest) SetCompanyId(_companyId int64) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusTopologyGetallRequest) GetCompanyId() int64 {
    return r._companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusTopologyGetallRequest) SetSystemId(_systemId string) error {
    r._systemId = _systemId
    r.Set("system_id", _systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusTopologyGetallRequest) GetSystemId() string {
    return r._systemId
}
