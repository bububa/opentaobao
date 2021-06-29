package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单项/加项包信息同步 APIResponse
alibaba.alihealth.examination.items.publish

体检机构对接_单项/加项包信息发布／更新
*/
type AlibabaAlihealthExaminationItemsPublishAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationItemsPublishResponse
}

type AlibabaAlihealthExaminationItemsPublishResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_items_publish_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
