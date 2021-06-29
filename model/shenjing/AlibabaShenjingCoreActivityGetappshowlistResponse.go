package shenjing

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取神鲸活动列表 APIResponse
alibaba.shenjing.core.activity.getappshowlist

获取神鲸活动列表
*/
type AlibabaShenjingCoreActivityGetappshowlistAPIResponse struct {
    model.CommonResponse
    AlibabaShenjingCoreActivityGetappshowlistResponse
}

type AlibabaShenjingCoreActivityGetappshowlistResponse struct {
    XMLName xml.Name `xml:"alibaba_shenjing_core_activity_getappshowlist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回的结果对象
    
    Result   *PageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
