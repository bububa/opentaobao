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
    campusId   int64
    // 公司id
    companyId   int64
    // 系统id
    systemId   string
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
func (r *AlibabaCampusTopologyGetallRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusTopologyGetallRequest) GetCampusId() int64 {
    return r.campusId
}
// CompanyId Setter
// 公司id
func (r *AlibabaCampusTopologyGetallRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusTopologyGetallRequest) GetCompanyId() int64 {
    return r.companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusTopologyGetallRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusTopologyGetallRequest) GetSystemId() string {
    return r.systemId
}
