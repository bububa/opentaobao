package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务商服务产品查询 APIResponse
tmall.mallitemcenter.serviceproduct.query

查询天猫服务的服务商发布的服务产品
*/
type TmallMallitemcenterServiceproductQueryAPIResponse struct {
    model.CommonResponse
    TmallMallitemcenterServiceproductQueryResponse
}

type TmallMallitemcenterServiceproductQueryResponse struct {
    XMLName xml.Name `xml:"tmall_mallitemcenter_serviceproduct_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TmallMallitemcenterServiceproductQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
