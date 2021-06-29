package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据ids和类型查询空间列表 APIResponse
alibaba.campus.space.getbyids

根据ids和类型查询空间列表
*/
type AlibabaCampusSpaceGetbyidsAPIResponse struct {
    model.CommonResponse
    AlibabaCampusSpaceGetbyidsResponse
}

type AlibabaCampusSpaceGetbyidsResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_space_getbyids_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 外卖订单查询结果
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
