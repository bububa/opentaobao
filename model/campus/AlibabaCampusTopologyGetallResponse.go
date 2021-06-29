package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取管理园区的规则拓扑接口 APIResponse
alibaba.campus.topology.getall

获取所属园区的所有规则拓扑图
*/
type AlibabaCampusTopologyGetallAPIResponse struct {
    model.CommonResponse
    AlibabaCampusTopologyGetallResponse
}

type AlibabaCampusTopologyGetallResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_topology_getall_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
