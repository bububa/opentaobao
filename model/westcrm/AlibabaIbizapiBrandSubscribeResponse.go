package westcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关注品牌号 APIResponse
alibaba.ibizapi.brand.subscribe

关注品牌号服务
*/
type AlibabaIbizapiBrandSubscribeAPIResponse struct {
    model.CommonResponse
    AlibabaIbizapiBrandSubscribeResponse
}

type AlibabaIbizapiBrandSubscribeResponse struct {
    XMLName xml.Name `xml:"alibaba_ibizapi_brand_subscribe_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否关注
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
