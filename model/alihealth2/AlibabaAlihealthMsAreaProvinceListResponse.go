package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗预约省份列表查询 APIResponse
alibaba.alihealth.ms.area.province.list

微信小程序疫苗预约省份列表查询
*/
type AlibabaAlihealthMsAreaProvinceListAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMsAreaProvinceListResponse
}

type AlibabaAlihealthMsAreaProvinceListResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_ms_area_province_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
