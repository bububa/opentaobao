package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取管理园区的规则拓扑接口 APIRequest
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

func NewAlibabaCampusTopologyGetallRequest() *AlibabaCampusTopologyGetallRequest{
    return &AlibabaCampusTopologyGetallRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusTopologyGetallRequest) GetApiMethodName() string {
    return "alibaba.campus.topology.getall"
}

func (r AlibabaCampusTopologyGetallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusTopologyGetallRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaCampusTopologyGetallRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaCampusTopologyGetallRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaCampusTopologyGetallRequest) GetCompanyId() int64 {
    return r.companyId
}

func (r *AlibabaCampusTopologyGetallRequest) SetSystemId(systemId string) error {
    r.systemId = systemId
    r.Set("system_id", systemId)
    return nil
}

func (r AlibabaCampusTopologyGetallRequest) GetSystemId() string {
    return r.systemId
}

