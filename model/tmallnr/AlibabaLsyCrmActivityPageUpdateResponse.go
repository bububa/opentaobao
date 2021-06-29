package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV活动页面创建修改 APIResponse
alibaba.lsy.crm.activity.page.update

ISV活动页面创建修改
*/
type AlibabaLsyCrmActivityPageUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaLsyCrmActivityPageUpdateResponse
}

type AlibabaLsyCrmActivityPageUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_lsy_crm_activity_page_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AlibabaLsyCrmActivityPageUpdateResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
