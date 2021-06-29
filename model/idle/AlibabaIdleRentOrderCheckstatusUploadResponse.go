package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传验收结果 APIResponse
alibaba.idle.rent.order.checkstatus.upload

上传验收结果
*/
type AlibabaIdleRentOrderCheckstatusUploadAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRentOrderCheckstatusUploadResponse
}

type AlibabaIdleRentOrderCheckstatusUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_rent_order_checkstatus_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
