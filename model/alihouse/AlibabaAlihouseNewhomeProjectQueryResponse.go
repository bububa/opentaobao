package alihouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询楼盘相关信息 APIResponse
alibaba.alihouse.newhome.project.query

根据outerid查询楼盘相关信息
*/
type AlibabaAlihouseNewhomeProjectQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAlihouseNewhomeProjectQueryResponse
}

type AlibabaAlihouseNewhomeProjectQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihouseNewhomeProjectQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
