package fundplatform

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询制卡商制卡进度 APIResponse
alibaba.fundplatform.cardorder.status.query

当通知制卡商进行制卡后，其制卡流程会比较长，若长时间未反馈当前制卡进展，则需要使用该接口来向制卡商发起进度查询。
*/
type AlibabaFundplatformCardorderStatusQueryAPIResponse struct {
    model.CommonResponse
    AlibabaFundplatformCardorderStatusQueryResponse
}

type AlibabaFundplatformCardorderStatusQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_fundplatform_cardorder_status_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结构体
    
    Response   *AlibabaFundplatformCardorderStatusQueryStruct `json:"response,omitempty" xml:"response,omitempty"`

    
}
