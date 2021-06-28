package wlbimports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取所有服务列表 APIResponse
taobao.wlb.imports.resource.get

一般进口TOP接口，获取所有服务列表。
*/
type TaobaoWlbImportsResourceGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbImportsResourceGetResponse
}

type TaobaoWlbImportsResourceGetResponse struct {
    XMLName xml.Name `xml:"wlb_imports_resource_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 一般进口所有服务商列表
    
    Resources   []ResourceResult `json:"resources,omitempty" xml:"resources>resource_result,omitempty"`
    
    
}
