package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
私域导购数据回流接口 APIResponse
alibaba.lsy.crm.activity.data.update

私域导购数据回流接口
*/
type AlibabaLsyCrmActivityDataUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaLsyCrmActivityDataUpdateResponse
}

type AlibabaLsyCrmActivityDataUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_lsy_crm_activity_data_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaLsyCrmActivityDataUpdateResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
