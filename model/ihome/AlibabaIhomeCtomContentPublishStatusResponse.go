package ihome

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
实拍图发布审核状态查询API APIResponse
alibaba.ihome.ctom.content.publish.status

实拍图发布审核状态查询API
*/
type AlibabaIhomeCtomContentPublishStatusAPIResponse struct {
    model.CommonResponse
    AlibabaIhomeCtomContentPublishStatusResponse
}

type AlibabaIhomeCtomContentPublishStatusResponse struct {
    XMLName xml.Name `xml:"alibaba_ihome_ctom_content_publish_status_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 根据站点名称查询产品
    
    Result   *AlibabaIhomeCtomContentPublishStatusApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
