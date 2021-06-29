package shenjing

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取神鲸活动列表 API返回值 
alibaba.shenjing.core.activity.getappshowlist

获取神鲸活动列表
*/
type AlibabaShenjingCoreActivityGetappshowlistAPIResponse struct {
    model.CommonResponse
    AlibabaShenjingCoreActivityGetappshowlistResponse
}

// 获取神鲸活动列表 成功返回结果
type AlibabaShenjingCoreActivityGetappshowlistResponse struct {
    XMLName xml.Name `xml:"alibaba_shenjing_core_activity_getappshowlist_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回的结果对象
    Result   *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}
