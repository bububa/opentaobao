package normalvisa

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
签证批量申请人送签接口 APIResponse
alitrip.travel.visa.sign.send

签证批量申请人送签接口，用于商家批量送签。
*/
type AlitripTravelVisaSignSendAPIResponse struct {
    model.CommonResponse
    AlitripTravelVisaSignSendResponse
}

type AlitripTravelVisaSignSendResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_visa_sign_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 批次信息
    
    BatchInfos   []BatchInfo `json:"batch_infos,omitempty" xml:"batch_infos>batch_info,omitempty"`
    
    
    // 失败信息
    
    FailInfos   []SendSignFailInfo `json:"fail_infos,omitempty" xml:"fail_infos>send_sign_fail_info,omitempty"`
    
    
}
