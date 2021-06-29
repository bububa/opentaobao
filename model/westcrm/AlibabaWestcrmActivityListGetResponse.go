package westcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取活动列表接口 APIResponse
alibaba.westcrm.activity.list.get

获取活动列表提供给友盟&互动屏
*/
type AlibabaWestcrmActivityListGetAPIResponse struct {
    model.CommonResponse
    AlibabaWestcrmActivityListGetResponse
}

type AlibabaWestcrmActivityListGetResponse struct {
    XMLName xml.Name `xml:"alibaba_westcrm_activity_list_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象封装
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
