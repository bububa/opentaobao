package aedropshiper

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询物流追踪信息 APIResponse
aliexpress.logistics.ds.trackinginfo.query

Dropshipper查询物流追踪信息
*/
type AliexpressLogisticsDsTrackinginfoQueryAPIResponse struct {
    model.CommonResponse
    AliexpressLogisticsDsTrackinginfoQueryResponse
}

type AliexpressLogisticsDsTrackinginfoQueryResponse struct {
    XMLName xml.Name `xml:"aliexpress_logistics_ds_trackinginfo_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 追踪详细信息列表
    
    Details   []Details `json:"details,omitempty" xml:"details>details,omitempty"`
    
    
    // 追踪网址
    
    OfficialWebsite   string `json:"official_website,omitempty" xml:"official_website,omitempty"`

    
    // error description
    
    ErrorDesc   string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`

    
    // success
    
    ResultSuccess   bool `json:"result_success,omitempty" xml:"result_success,omitempty"`

    
}
