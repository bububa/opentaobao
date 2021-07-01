package westcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取活动列表接口 API返回值 
alibaba.westcrm.activity.list.get

获取活动列表提供给友盟&互动屏
*/
type AlibabaWestcrmActivityListGetAPIResponse struct {
    model.CommonResponse
    AlibabaWestcrmActivityListGetAPIResponseModel
}

// 获取活动列表接口 成功返回结果
type AlibabaWestcrmActivityListGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_westcrm_activity_list_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回对象封装
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
