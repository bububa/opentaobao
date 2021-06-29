package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】开放平台对外页面跳转 APIResponse
alitrip.btrip.openplatform.address.get

获取类目预定页跳转地址
*/
type AlitripBtripOpenplatformAddressGetAPIResponse struct {
    model.CommonResponse
    AlitripBtripOpenplatformAddressGetResponse
}

type AlitripBtripOpenplatformAddressGetResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_openplatform_address_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
